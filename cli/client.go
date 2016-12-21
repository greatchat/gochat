package cli

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/pkg/errors"

	"github.com/greatchat/gochat/cache"
	"github.com/greatchat/gochat/filter"
	chat "github.com/greatchat/gochat/transport"
)

// User is the current user using the cli
type User struct {
	Username      string `json:"username"`
	SecretKeyRing string `json:"secretKeyRing"`
	PublicKeyRing string `json:"publicKeyRing"`
	Passphrase    string `json:"passphrase"`
}

// encode / decode here : https://gist.github.com/stuart-warren/93750a142d3de4e8fdd2

// Start starts an interactive cli.
func Start(ctx context.Context, out io.Writer, in io.Reader, client chat.Client) (func(), error) {
	stopChan := make(chan struct{})

	u, err := login()
	if err != nil {
		log.Fatal("Unable to login the current user")
	}
	log.WithField("email", u.Username).Info("Logged in user")

	var (
		producerName string
		consumerName string
	)

	prompt(&producerName, "Enter the room name:")
	//prompt(&consumerName, "Consumer Name:")
	consumerName = producerName
	log.WithFields(log.Fields{
		"producerName": producerName,
		"consumerName": consumerName,
	}).Info("Connecting...")

	loginMsg := fmt.Sprintf("[Logged in as %s]\n", u.Username)
	out.Write([]byte(loginMsg))

	receive, err := client.Consumer(consumerName)
	if err != nil {
		log.WithError(err).Fatal("failed to set up consumer")
	}

	send, err := client.Producer(producerName)
	if err != nil {
		log.WithError(err).Fatal("failed to set up producer")
	}

	reader := bufio.NewReader(in)
	defState := fmt.Sprintf("\033[0;33m%s [%s]\033[0m > ", u.Username, producerName)

	//TODO have a look at pipes to replace the old line http://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string

	go func() {
		for {
			msg := <-receive
			log.WithFields(log.Fields{
				"body":      msg.Body,
				"author":    msg.Author,
				"timestamp": msg.Timestamp.Format(time.RFC3339),
			}).Debug("Received message")

			// discard the message if the author and the receiver are the same
			if msg.Author != u.Username {
				outPrint := "\r\033[0;32m%s@%s %s\033[0m > %s"
				out.Write([]byte(fmt.Sprintf(outPrint, msg.Author, producerName, msg.Timestamp.Format(time.Kitchen), msg.Body)))
				out.Write([]byte(defState))

			}

		}
	}()

	go func() {
		for {
			select {
			case <-stopChan:
				out.Write([]byte("\nStopping now...\n"))
				return
			default:
				out.Write([]byte(defState))
				text, _ := reader.ReadString('\n')
				log.Debugf("Sending %s", text)

				// Filter here.
				filter.Apply(&text, filter.Dev, filter.Division, filter.UK)
				msg := chat.Message{
					Author:    u.Username,
					Body:      text,
					Timestamp: time.Now(),
				}

				send <- msg
				outPrint := "\x1b[1A\x1b[2K\033[0;31m%s@%s %s \033[0m > %s"
				out.Write([]byte(fmt.Sprintf(outPrint, u.Username, producerName, msg.Timestamp.Format(time.Kitchen), text)))

				log.Debug("Message sent")
			}
		}
	}()

	stop := func() {
		stopChan <- struct{}{}
	}

	return stop, nil
}

func write(out io.Writer) {

}

// StartTest is just a test
func StartTest() {
	login()
}

// ask for login
func login() (*User, error) {

	cacheFile, err := cache.FilePath(".credentials", url.QueryEscape("go-chat.sainbsurys.co.uk.json"))
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}

	u, err := userFromFile(cacheFile)
	if err != nil {
		log.Warn("User not set: authenticating")
		u := &User{}
		promptForCredentials(u)
		err = authenticate(u)

		if err != nil {
			log.Fatal("Failed to authenticate to the Chat  service")
		}

		saveUser(cacheFile, u)
		return u, nil
	}

	return u, nil

}

func authenticate(u *User) error {
	// TODO perform some checks on GPG key, talk to a server...
	return nil
}

func promptForCredentials(u *User) {
	prompt(&u.Username, "Enter your Username:")
	prompt(&u.PublicKeyRing, "Enter the location of your public key ring:")
	prompt(&u.SecretKeyRing, "Enter the location of your secret key ring:")
	prompt(&u.Passphrase, "Enter the passphrase (or leave empty):")
}

func prompt(s *string, msg string) {
	fmt.Println(msg)
	if _, err := fmt.Scan(s); err != nil {
		log.Fatalf("Unable to read the attribute %v", err)
	}
}

func userFromFile(file string) (*User, error) {

	f, err := cache.Open(file)

	if err != nil {
		return nil, errors.Wrapf(err, "The cache file %s was not found", file)
	}
	u := &User{}
	err = json.NewDecoder(f).Decode(u)
	defer f.Close()
	log.Debug("User loaded from cache")
	return u, err

}

// saveToken uses a file path to create a file and store the
// token in it.
func saveUser(file string, u *User) {
	log.Debugf("Saving credential file to: %s\n", file)
	f, err := cache.OpenOrCreate(file)

	if err != nil {
		log.Fatalf("Unable to cache user: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(u)
}

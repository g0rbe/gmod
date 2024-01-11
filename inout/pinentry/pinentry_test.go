package pinentry

import (
	"github.com/g0rbe/gmod/inout/logz"
)

func ExampleRun() {

	result, err := Run(
		CommandSetTimeout(60),
		CommandSetDescription("Description"),
		CommandSetPrompt("Prompt"),
		CommandSetTitle("Title"),
		CommandSetOk("Ok"),
		CommandSetCancel("Cancel"),
		CommandConfirm(),
		CommandGetPin(),
		CommandBye())

	if err != nil {
		logz.Errorf("Failed to run: %s\n", err)
		return
	}

	logz.Debugf("Result: %s\n", result)
}

func ExampleMessage() {

	err := Message("TestMessage Prompt", "TestMessage Description", "TestMessage")

	if err != nil {
		logz.Errorf("Failed to run: %s\n", err)
	}

}

func ExampleConfirm() {

	confirmed, err := Confirm("TestMessage Prompt", "TestMessage Description", "Allow", "Deny")

	if err != nil {
		logz.Errorf("Failed to run: %s\n", err)
		return
	}

	if confirmed {
		// Confirmed
	}

}

func ExamplePassword() {

	pass, err := Password("Passphrase:", "please enter the passphrase for the ssh key ...")
	if err != nil {
		logz.Errorf("Failed to run: %s\n", err)
		return
	}

	logz.Debugf("Pass: %s\n", pass)
}

func ExamplePasswordRepeat() {

	pass, err := PasswordRepeat("Passphrase:", "Enter passphrase")
	if err != nil {
		logz.Errorf("Failed to run: %s\n", err)
		return
	}

	logz.Debugf("Pass: %s\n", pass)
}

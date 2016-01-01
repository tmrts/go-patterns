package main

import "os"

type Arguments struct {
	UID         int
	GID         int
	Flags       int
	Contents    string
	Permissions os.FileMode
}

type Argument func(*Arguments)

func UID(userID int) Argument {
	return func(args *Arguments) {
		args.UID = userID
	}
}

func GID(groupID int) Argument {
	return func(args *Arguments) {
		args.GID = groupID
	}
}

func Contents(c string) Argument {
	return func(args *Arguments) {
		args.Contents = c
	}
}

func Permissions(perms os.FileMode) Argument {
	return func(args *Arguments) {
		args.Permissions = perms
	}
}

func New(filepath string, setters ...Argument) error {
	// Default Arguments
	args := &Arguments{
		UID:         os.Getuid(),
		GID:         os.Getgid(),
		Contents:    "",
		Permissions: 0666,
		Flags:       os.O_CREATE | os.O_EXCL | os.O_WRONLY,
	}

	for _, setter := range setters {
		setter(args)
	}

	f, err := os.OpenFile(filepath, args.Flags, args.Permissions)
	if err != nil {
		return err
	} else {
		defer f.Close()
	}

	if _, err := f.WriteString(args.Contents); err != nil {
		return err
	}

	return f.Chown(args.UID, args.GID)
}

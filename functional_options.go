package main

import "os"

type arguments struct {
	UID         int
	GID         int
	Flags       int
	Contents    string
	Permissions os.FileMode
}

type argument func(*arguments)

func UID(userID int) argument {
	return func(args *arguments) {
		args.UID = userID
	}
}

func GID(groupID int) argument {
	return func(args *arguments) {
		args.GID = groupID
	}
}

func Contents(c string) argument {
	return func(args *arguments) {
		args.Contents = c
	}
}

func Permissions(perms os.FileMode) argument {
	return func(args *arguments) {
		args.Permissions = perms
	}
}

func New(filepath string, setters ...argument) error {
	// Default arguments
	args := &arguments{
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

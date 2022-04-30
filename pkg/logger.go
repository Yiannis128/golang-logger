package pkg

import (
	"log"
	"os"
)

type Logger struct {
	verbosity uint
}

func (l *Logger) SetVerbosity(level uint) {
	l.verbosity = level
}

func (l *Logger) Println(level uint, v ...any) {
	if l.verbosity >= level {
		log.Println(v...)
	}
}

func (l *Logger) Print(level uint, v ...any) {
	if l.verbosity >= level {
		log.Print(v...)
	}
}

func (l *Logger) Printf(level uint, f string, v ...any) {
	if l.verbosity >= level {
		log.Printf(f, v...)
	}
}

func (l *Logger) Panic(level uint, v ...any) {
	if l.verbosity >= level {
		log.Panic(v...)
	} else {
		panic(v)
	}
}

func (l *Logger) Panicln(level uint, v ...any) {
	if l.verbosity >= level {
		log.Panicln(v...)
	} else {
		panic(v)
	}
}

func (l *Logger) Panicf(level uint, f string, v ...any) {
	if l.verbosity >= level {
		log.Panicf(f, v...)
	} else {
		panic(v)
	}
}

func (l *Logger) Fatal(level uint, v ...any) {
	if l.verbosity >= level {
		log.Fatal(v...)
	} else {
		os.Exit(1)
	}
}

func (l *Logger) Fatalln(level uint, v ...any) {
	if l.verbosity >= level {
		log.Fatalln(v...)
	} else {
		os.Exit(1)
	}
}

func (l *Logger) Fatalf(level uint, f string, v ...any) {
	if l.verbosity >= level {
		log.Fatalf(f, v...)
	} else {
		os.Exit(1)
	}
}

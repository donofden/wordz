# ıllıllı [ WORDZ ] ıllıllı
A CLI Dictionary Application, where you type a word and get its meaning from Oxford Dictionaries. Based on the Oxford Dictionaries API. Wordz lookup should give you meaning, pronunciation, sample sentances, synonyms and antonyms (if applicable).

![Full screen](wordz-logo.png)

## Configuration

To run/use this application, you need `API Key` and `API_ID` from `Oxford Dictionaries` which you can create for free. [Follow the link here](https://developer.oxforddictionaries.com/login)

You will need the following environment variables defining:

```config
OXFORD_APPLICATION_ID="Application_ID"
OXFORD_APPLICATION_KEY="Application_Keys"
OXFORD_VOICE_ACTIVATE=1
```
Note: `OXFORD_VOICE_ACTIVATE` will pronunce the word, set the value to `1` in environment variable to activate it, `0` to disable the voice pronunciation.


## How to Install and Run

- Run the following command to know the availabe options.

Install

```Makefile
make install
```

To Build

```Makefile
make build
```

To Run and install the binary

```Makefile
make build-project
```

Other Options Availabe:

```Makefile
### Welcome
#
# __          __           _     
# \ \        / /          | |    
#  \ \  /\  / /__  _ __ __| |____
#   \ \/  \/ / _ \| '__/ _` |_  /
#    \  /\  / (_) | | | (_| |/ / 
#     \/  \/ \___/|_|  \__,_/___|
#
#
### Installation
#
# $ make install
#
### Targets
 Choose a command run in wordz:
build-project                  Build the CLI in PROJECT bin folder
build                          Build the CLI in Project Folder
clean                          Clean the Project
get                            Get Go Dependencies
install                        Install "go install GOFILES"
run                            Run the Project
start                          Start the application "bin/PROJECT"
stop                           Stop the application  "bin/PROJECT"
```

## Usage

Once the binary is built, lets make it easy for us to use it. Create a aliases for the application.

**Open zsh configuration file:** `sudo vim ~/.zshrc`

**Add this to the file:** `alias wz='wordz find'`

**Apply the changes:** `source ~/.zshrc` or you can restart your terminal session.

> ## *Let’s begin*

```cmd
~ » wz excellent

Please wait while we search for the meaning...

 Word:  excellent

Loading.. .. .. 🌏     ¯\\_(ツ)_/¯
 ```
# Twitch Raid Finder

Finding someone to raid after your stream can be a time consuming task,
especially if you are not sure who you are looking for.

Now, thanks to twitch raid finder, you can easily find someone to raid without getting distracted from your last game

# Installation

Head to the [releases page](https://github.com/mrTomatolegit/raid-finder/releases) and download the latest version of the program depending on your operating system (tested for windows)

Extract the zip contents into a new folder

# Usage

> We recommend using it once before starting your stream to avoid having it reauthorize in the middle of your session

To use the program, double click it. It will open a browser window where you will log in with twitch

Once you authorize the app, the browser tab will close and the program will start finding people to raid

The users found by the app will be displayed in the terminal

# Raidlist and NoRaidlist

These lists allow the app to search for/ignore specific streamers.

The `raidlist.txt` file will make the app check the livestream of each username

The `noraidlist.txt` file will make the app ignore the streams from the raidlist, your followed channels, and searched channels

## Format

Every list uses the same format, each username to search for must be on a new line

**Example:**
```
TimothySoup
Markiplier
Jacksepticeye
```

# Updating

If you'd like to update the program to a newer version, take the new raidfinder file and move it into the same folder as the old one, this will keep your old raidlist/noraidlist and cached token

# Contribution

If you'd like to contribute to the project, please fork it and create a pull request

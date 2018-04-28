# 觀察貓

### 觀察檔案異動執行相關動做

```sh
NAME:
   OserverdCat - A new cli application

USAGE:
   observedcat [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR:
   Ken

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --file value         ObservedFile (default: "ObservedFile") [$OBSERVED_FILE]
   --eventall value     EventAllExec (default: "echo 'EventAll'") [$EVENT_ALL_EXEC]
   --eventcreate value  EventCreateExec (default: "echo 'EventCreateExec'") [$EVENT_CREATE_EXEC]
   --eventwrite value   EventWriteExec (default: "echo 'EventWriteExec'") [$EVENT_WRITE_EXEC]
   --eventremove value  EventRemoveExec (default: "echo 'EventRemoveExec'") [$EVENT_REMOVE_EXEC]
   --eventrename value  EventRenameExec (default: "echo 'EventRenameExec'") [$EVENT_RENAME_EXEC]
   --eventchmod value   EventChmodExec (default: "echo 'EventChmodExec'") [$EVENT_CHMOD_EXEC]
   --help, -h           show help
   --version, -v        print the version
```

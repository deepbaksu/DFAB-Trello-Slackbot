# Trello Task Helper CLI

1. Create environment variables

```sh
export TRELLO_APP_KEY=your-app-key-from-trello
export TRELLO_TOKEN=your-trello-token
export TRELLO_BOARD_ID=your-trello-board-id
export TRELLO_USERNAME=your-trello-username
```

2. Run the command

```sh
go run main.go
```

Then you will see

```
To Do
=====
- Set up a CD for the pipeline

Doing
=====
- Finish TFJS chapter
- Set up a meeting

Done
====
- Answer emails
- Update IDs
```

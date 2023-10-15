# gotrack
A simple CLI for tracking time, written in Go.

> [!WARNING]  
> Not production ready—do not use!

## Track time
- `gotrack [start time] [end time] [break duration] [date]`
- All arguments except `start time` are optional.
- Examples:
    - `gotrack 09:00 17:00 0:30 2023-08-22` -- Tracks time for 22nd August 2023. Start time: 9:00, end time 17:00, break duration 30 minutes. Results in a total duration of 7h30m.
    - `gotrack 9 17 1` -- Tracks time for today. Start time 9:00, end time 17:00, break duration 1 hour. Results in a total duration of 7h.
    - `gotrack 9 17:30` -- Tracks time for today. Start time 9:00, end time 17:30, no break. Results in a total duration of 8h30m.
    - `gotrack 9` -- Tracks time for today. Start time 9:00, end time `time.Now()`, no break.

## Print tracked time
- `gotrack print [date]`
- All arguments are optional.
- Examples:
    - `gotrack print 2023-08-22 2023-08-30` -- Prints tracked times from 22nd August 2023 to 30th August 2023 (both inclusive).
    - `gotrack print 2023-08-22` -- Prints tracked time for 22nd August 2023.
    - `gotrack print` -- Prints tracked time for today.

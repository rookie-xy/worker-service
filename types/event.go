/*
 * Copyright (C) 2016 Meng Shi
 */

package types

type Event struct {
    name    string
    opcode  int
    gid     uint64

    notice  chan *Event
}

func NewEvent() *Event {
    return &Event{
        notice:make(chan *Event),
    }
}

func (e *Event) SetName(name string) int {
    if name == "" {
        return Error
    }

    e.name = name

    return Ok
}

func (e *Event) GetName() string {
    return e.name
}

func (e *Event) SetOpcode(opcode int) int {
    e.opcode = opcode
    return Ok
}

func (e *Event) GetOpcode() int {
    return e.opcode
}

func (e *Event) GetNotice() chan *Event {
    return e.notice
}

func (e *Event) SetNotice() {
    e.notice <- e
}

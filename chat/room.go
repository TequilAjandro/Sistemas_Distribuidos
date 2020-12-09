package main

type room struct {
	Name     string
	Members  []*client
	Messages []string
}

func (r *room) broadcast(sender *client, msg string) {
	for _, m := range r.Members {
		if sender.nick != m.nick {
			m.msg(msg)
		}
	}
	r.Messages = append(r.Messages, (sender.conn.RemoteAddr().String() + " " + sender.nick + ": " + msg))
}

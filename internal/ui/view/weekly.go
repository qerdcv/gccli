package view

import (
	"errors"
	"fmt"
	"github.com/jroimartin/gocui"
	"google.golang.org/api/calendar/v3"
	"strconv"
	"strings"
	"time"
)

const (
	dayInWeek = 7
)

type day struct {
	fullName  string
	shortName string

	date time.Time
}

var days = [dayInWeek]day{
	{shortName: "SUN"},
	{shortName: "MON"},
	{shortName: "TUE"},
	{shortName: "WED"},
	{shortName: "THU"},
	{shortName: "FRI"},
	{shortName: "SAT"},
}

type WeeklyView struct {
	days          [dayInWeek]day
	currentDayIdx int
}

func NewWeeklyView() *WeeklyView {
	return &WeeklyView{
		days: days,
	}
}

func (wv *WeeklyView) Layout(g *gocui.Gui, events *calendar.Events) error {
	maxX, maxY := g.Size()
	daySize := maxX / dayInWeek

	currentDay := time.Now().Day()
	currentDayOfWeek := int(time.Now().Weekday())
	for i := 0; i < dayInWeek; i++ {
		d := wv.days[i]
		v, err := g.SetView(fmt.Sprintf("day-%d", i), i*daySize, 0, (i+1)*daySize, maxY)
		if err != nil && !errors.Is(err, gocui.ErrUnknownView) {
			return fmt.Errorf("gui set view: %w", err)
		}

		v.Clear()

		if i == wv.currentDayIdx {
			v.BgColor = gocui.ColorBlue
		}

		fmt.Fprint(v, strings.Repeat(" ", daySize/2-5))

		if i == currentDayOfWeek {
			fmt.Fprintf(v, "\033[41m%s\033[0m", d.shortName)
		} else {
			fmt.Fprint(v, d.shortName)
		}

		fmt.Fprintln(v, " "+strconv.Itoa((currentDay-currentDayOfWeek)+i))

		for _, e := range events.Items {
			t, err := time.Parse(time.RFC3339, e.Start.DateTime)
			if err != nil {
				return fmt.Errorf("parse event time: %w", err)
			}

			if int(t.Weekday()) == i {
				fmt.Fprintln(v, e.Summary)
			}
		}
	}

	return nil
}

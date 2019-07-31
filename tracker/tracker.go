package tracker

import (
	"fmt"

	"gopkg.in/salsita/go-pivotaltracker.v2/v5/pivotal"
)

type Tracker struct {
	c         *pivotal.Client
	projectID int
}

type ConfigOpts struct {
	Token     string
	ProjectID int
}

func NewTracker(opts ConfigOpts) Tracker {
	client := pivotal.NewClient(opts.Token)
	client.SetUserAgent("trackertix")

	return Tracker{
		c:         client,
		projectID: opts.ProjectID,
	}
}

func (t *Tracker) SetupProject() error {
	var (
		stories     []*pivotal.Story
		todayMarker bool
		soonMarker  bool
		inboxMarker bool
	)
	path := fmt.Sprintf("projects/%d/stories?with_story_type=release", t.projectID)
	req, err := t.c.NewRequest("GET", path, nil)
	if err != nil {
		return err
	}

	res, err := t.c.Do(req, &stories)
	if err != nil {
		return err
	}
	if res.StatusCode > 299 {
		return fmt.Errorf("received response code of %d", res.StatusCode)
	}

	for i := range stories {
		if stories[i].Name == "⬆️ Due Today ⬆️" {
			todayMarker = true
		}
		if stories[i].Name == "⬆️ Due Soon ⬆️" {
			soonMarker = true
		}
		if stories[i].Name == "⬆️ Inbox ⬆️" {
			inboxMarker = true
		}
	}

	storyReq := &pivotal.StoryRequest{
		Type: "release",
	}

	if !todayMarker {
		storyReq.Name = "⬆️ Due Today ⬆️"
		storyReq.Description = "Stories above this marker should completed today."
		storyReq.State = "unstarted"
		_, res, err := t.c.Stories.Create(t.projectID, storyReq)
		if err != nil {
			return err
		}
		if res.StatusCode > 299 {
			return fmt.Errorf("received response code of %d", res.StatusCode)
		}
	}

	if !soonMarker {
		storyReq.Name = "⬆️ Due Soon ⬆️"
		storyReq.Description = "Stories above this marker should completed soon."
		storyReq.State = "unstarted"
		_, res, err := t.c.Stories.Create(t.projectID, storyReq)
		if err != nil {
			return err
		}
		if res.StatusCode > 299 {
			return fmt.Errorf("received response code of %d", res.StatusCode)
		}
	}

	if !inboxMarker {
		storyReq.Name = "⬆️ Inbox ⬆️"
		storyReq.Description = "Stories above this marker need to be triaged."
		storyReq.State = "unscheduled"
		_, res, err := t.c.Stories.Create(t.projectID, storyReq)
		if err != nil {
			return err
		}
		if res.StatusCode > 299 {
			return fmt.Errorf("received response code of %d", res.StatusCode)
		}
	}

	return nil
}

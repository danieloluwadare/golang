package main

import "fmt"

type event struct {
	ID          string
	Title       string
	Description string
	mapofE      map[int]event
}

type eContainsDeriviedEasprop struct {
	ID          string
	Title       string
	Description string
	mapofE      mapOfEvents
}

type allEvents []event

type mapOfEvents map[int]event

// -------------------------------------
// This type below mapintToListOfEvent == can be used interchangebly
// because they are achieving the same function
// both represent Java Map<Integer List<Event>>
type mapintToListOfEvent map[int][]event

type mapintToAllEvents map[int]allEvents

// --------------------------------------

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

func testTypePlayGround() {
	map1 := make(map[int]event)
	event1 := event{
		ID:          "1",
		Title:       "event1",
		Description: "event1",
	}
	event2 := event{
		ID:          "2",
		Title:       "Golang2",
		Description: "2",
	}
	event4 := event{
		ID:          "4",
		Title:       "Golang4",
		Description: "4",
	}

	map1[1] = event1
	map1[2] = event2

	event3 := event{
		ID:          "1",
		Title:       " Golang1",
		Description: "checkitout",
		mapofE:      map1,
	}

	map2 := map[int]event{
		1: event{
			ID:          "1",
			Title:       "evenTContainD1 ",
			Description: "evenTContainD1 out",
		},
	}

	var mofEventsOfDerivedType mapOfEvents = map2

	evenTContainD1 := eContainsDeriviedEasprop{
		ID:          "1",
		Title:       " Golang1",
		Description: "checkitout",
		mapofE:      map1,
	}

	evenTContainD2 := eContainsDeriviedEasprop{
		ID:          "1",
		Title:       " Golang1",
		Description: "checkitout",
		mapofE:      mofEventsOfDerivedType,
	}

	fmt.Printf("event3 => %v \n", event3)
	fmt.Printf("eContainsDeriviedEasprop => %v \n", evenTContainD1)
	_ = evenTContainD2
	// fmt.Printf("eContainsDeriviedEasprop => %v \n", evenTContainD2)

	var moe mapOfEvents = evenTContainD1.mapofE
	moe[4] = event4
	fmt.Printf("eContainsDeriviedEasprop => %v \n", evenTContainD1)

	// fmt.Printf("event3 => %v and of type %T\n", event3, event3)
	// fmt.Printf("eContainsDeriviedEasprop => %v and of type %T\n", evenTContainD1, evenTContainD1)
	// fmt.Printf("eContainsDeriviedEasprop => %v and of type %T\n", evenTContainD2, evenTContainD2)

}

func testDerivedTypesOfmapOfEvents() {
	var sliceOfEvents = allEvents{
		{
			ID:          "1",
			Title:       "T1",
			Description: "desc",
		},
		{
			ID:          "2",
			Title:       "T1",
			Description: "desc",
		},
	}

	var ma = mapintToListOfEvent{
		1: sliceOfEvents,
	}

	var allevet1 allEvents = sliceOfEvents

	var ma2 mapintToAllEvents = mapintToAllEvents{
		1: allevet1,
	}

	fmt.Printf("event3 => %v and of type %T\n", ma, ma)
	fmt.Printf("event3 => %v and of type %T\n", ma2, ma2)

}

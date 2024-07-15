import { CalendarItemGroupType } from "@/types/CalendarItemGroup"
import { ScriptProps } from "next/script"
import { createContext, useReducer } from "react"

type CalendarItemGroupsAction = 
| {
    type: "addCalendarItemGroup"
    payload: CalendarItemGroupType
}
| {
    type: "updateCalendarItemGroup"
    payload: CalendarItemGroupType
}
| {
    type: "removeCalendarItemGroup"
    payload: {
        id: string
    }
}

const calendarItemGroupsReducer = (state: CalendarItemGroupType[], action: CalendarItemGroupsAction): CalendarItemGroupType[] => {
    switch (action.type) {
        case "addCalendarItemGroup":
            return [...state, action.payload]
        case "updateCalendarItemGroup":
            return state.map((calendarItemGroup: CalendarItemGroupType) => calendarItemGroup.id === action.payload.id ? action.payload : calendarItemGroup)
        case "removeCalendarItemGroup":
            return state.filter((calendarItemGroup: CalendarItemGroupType) => calendarItemGroup.id !== action.payload.id)
    }    
}

const defaultCalendarItemGroupsState = [
    {
        id: "cccccccc000",
        name: "チキンカレー",
        calendarItemIds: [
            "bbbbbbbbbbbbbbb000",
            "bbbbbbbbbbbbbbb002",
            "bbbbbbbbbbbbbbb003",
        ],
    },
]

function useCalendarItemGroupsReducer() {
    const [state, dispatch] = useReducer(calendarItemGroupsReducer, defaultCalendarItemGroupsState)
    return { state, dispatch }
}

const defaultCalendarItemGroups: ReturnType<typeof useCalendarItemGroupsReducer> = {
    state: defaultCalendarItemGroupsState,
    dispatch: () => {},
}

const CalendarItemGroupsContext = createContext(defaultCalendarItemGroups)
function CalendarItemGroupsProvider (props: ScriptProps) {
    return (
        <CalendarItemGroupsContext.Provider value={useCalendarItemGroupsReducer()}>
            {props.children}
        </CalendarItemGroupsContext.Provider>
    )
}

export { useCalendarItemGroupsReducer, defaultCalendarItemGroups, CalendarItemGroupsContext, CalendarItemGroupsProvider }

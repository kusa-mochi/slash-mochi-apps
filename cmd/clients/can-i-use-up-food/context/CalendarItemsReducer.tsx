import { MealTime } from "@/connect/can_i_use_up_food_pb"
import { CalendarItemType } from "@/types/CalendarItem"
import { ScriptProps } from "next/script"
import { createContext, useReducer } from "react"

type CalendarItemsAction = 
| {
    type: "addCalendarItem"
    payload: CalendarItemType
}
| {
    type: "updateCalendarItem"
    payload: CalendarItemType
}
| {
    type: "removeCalendarItem"
    payload: {
        id: string
    }
}

const calendarItemsReducer = (state: CalendarItemType[], action: CalendarItemsAction): CalendarItemType[] => {
    switch (action.type) {
        case "addCalendarItem":
            return [...state, action.payload]
        case "updateCalendarItem":
            return state.map((calendarItem: CalendarItemType) => calendarItem.id === action.payload.id ? action.payload : calendarItem)
        case "removeCalendarItem":
            return state.filter((calendarItem: CalendarItemType) => calendarItem.id !== action.payload.id)
    }    
}

const defaultCalendarItemsState = [
    {
        id: "bbbbbbbbbbbbbbb000",
        foodId: "aaaaaaaaaa000",
        date: {
            year: 2024,
            month: 6,
            date: 26,
        },
        mealTime: MealTime.Dinner,
        amount: 1,
    },
    {
        id: "bbbbbbbbbbbbbbb001",
        foodId: "aaaaaaaaaa000",
        date: {
            year: 2024,
            month: 6,
            date: 28,
        },
        mealTime: MealTime.Lunch,
        amount: 1,
    },
    {
        id: "bbbbbbbbbbbbbbb002",
        foodId: "aaaaaaaaaa001",
        date: {
            year: 2024,
            month: 6,
            date: 26,
        },
        mealTime: MealTime.Dinner,
        amount: 2,
    },
    {
        id: "bbbbbbbbbbbbbbb003",
        foodId: "aaaaaaaaaa003",
        date: {
            year: 2024,
            month: 6,
            date: 26,
        },
        mealTime: MealTime.Dinner,
        amount: 1,
    },
    {
        id: "bbbbbbbbbbbbbbb004",
        foodId: "aaaaaaaaaa004",
        date: {
            year: 2024,
            month: 7,
            date: 8,
        },
        mealTime: MealTime.Morning,
        amount: 100,
    },
    {
        id: "bbbbbbbbbbbbbbb005",
        foodId: "aaaaaaaaaa004",
        date: {
            year: 2024,
            month: 7,
            date: 9,
        },
        mealTime: MealTime.Morning,
        amount: 100,
    },
    {
        id: "bbbbbbbbbbbbbbb006",
        foodId: "aaaaaaaaaa004",
        date: {
            year: 2024,
            month: 7,
            date: 10,
        },
        mealTime: MealTime.Morning,
        amount: 100,
    },
    {
        id: "bbbbbbbbbbbbbbb007",
        foodId: "aaaaaaaaaa005",
        date: {
            year: 2024,
            month: 7,
            date: 5,
        },
        mealTime: MealTime.Morning,
        amount: 100,
    },
    {
        id: "bbbbbbbbbbbbbbb008",
        foodId: "aaaaaaaaaa005",
        date: {
            year: 2024,
            month: 7,
            date: 6,
        },
        mealTime: MealTime.Morning,
        amount: 100,
    },
    {
        id: "bbbbbbbbbbbbbbb009",
        foodId: "aaaaaaaaaa005",
        date: {
            year: 2024,
            month: 7,
            date: 7,
        },
        mealTime: MealTime.Morning,
        amount: 100,
    },
]

function useCalendarItemsReducer() {
    const [state, dispatch] = useReducer(calendarItemsReducer, defaultCalendarItemsState)
    return { state, dispatch }
}

const defaultCalendarItems: ReturnType<typeof useCalendarItemsReducer> = {
    state: defaultCalendarItemsState,
    dispatch: () => {},
}

const CalendarItemsContext = createContext(defaultCalendarItems)
function CalendarItemsProvider (props: ScriptProps) {
    return (
        <CalendarItemsContext.Provider value={useCalendarItemsReducer()}>
            {props.children}
        </CalendarItemsContext.Provider>
    )
}

export { useCalendarItemsReducer, defaultCalendarItems, CalendarItemsContext, CalendarItemsProvider }

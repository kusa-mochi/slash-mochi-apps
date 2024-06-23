import { AmountUnit, MealTime } from "@/connect/can_i_use_up_food_pb"
import { createContext, useReducer } from "react"

type DateType = {
    year: number    // ex: 2024
    month: number   // ex: 6 (means June)
    date: number    // ex: 31
}

type CalendarItemType = {
    id: string
    foodId: string
    date: DateType
    mealTime: MealTime
    amount: number
}

type FoodType = {
    id: string
    name: string
    totalAmount: number
    amountUnit: AmountUnit
    limitDate: DateType
}

type CalendarItemGroupType = {
    name: string
    calendarItemIds: string[]
}

type ProjectContextType = {
    name: string
    startDate: DateType
    endDate: DateType
    foods: FoodType[]
    calendarItems: CalendarItemType[]
    calendarItemGroups: CalendarItemGroupType[]
}

const initialProjectContextValue: ProjectContextType = {
    name: "テストプロジェクト",
    startDate: {
        year: 2024,
        month: 6,
        date: 25,
    },
    endDate: {
        year: 2024,
        month: 7,
        date: 10,
    },
    foods: [
        {
            id: "aaaaaaaaaa000",
            name: "鶏もも肉",
            totalAmount: 2,
            amountUnit: AmountUnit.PIECES,
            limitDate: {
                year: 2024,
                month: 6,
                date: 28,
            },
        },
        {
            id: "aaaaaaaaaa001",
            name: "玉ねぎ",
            totalAmount: 3,
            amountUnit: AmountUnit.PIECES,
            limitDate: {
                year: 2024,
                month: 7,
                date: 10,
            },
        },
        {
            id: "aaaaaaaaaa002",
            name: "卵",
            totalAmount: 8,
            amountUnit: AmountUnit.PIECES,
            limitDate: {
                year: 2024,
                month: 7,
                date: 4,
            },
        },
        {
            id: "aaaaaaaaaa003",
            name: "にんじん",
            totalAmount: 1,
            amountUnit: AmountUnit.PIECES,
            limitDate: {
                year: 2024,
                month: 6,
                date: 30,
            },
        },
        {
            id: "aaaaaaaaaa004",
            name: "コーンフレーク",
            totalAmount: 500,
            amountUnit: AmountUnit.GRAMS,
            limitDate: {
                year: 2024,
                month: 7,
                date: 10,
            },
        },
        {
            id: "aaaaaaaaaa005",
            name: "みかんジュース",
            totalAmount: 300,
            amountUnit: AmountUnit.MILLI_LITER,
            limitDate: {
                year: 2024,
                month: 7,
                date: 10,
            },
        },
    ],
    calendarItems: [
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
    ],
    calendarItemGroups: [
        {
            name: "チキンカレー",
            calendarItemIds: [
                "bbbbbbbbbbbbbbb000",
                "bbbbbbbbbbbbbbb002",
                "bbbbbbbbbbbbbbb003",
            ],
        },
    ],
}

type ProjectAction = 
| {
    type: "setProjectName"
    payload: {
        name: string
    }
}

const projectReducer = (state: ProjectContextType, action: ProjectAction): ProjectContextType => {
    const next: ProjectContextType = { ...state }

    switch (action.type) {
        case "setProjectName":
            next.name = state.name
            break
    }

    return next
}

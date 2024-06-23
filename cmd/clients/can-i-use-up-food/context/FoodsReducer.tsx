import { AmountUnit } from "@/connect/can_i_use_up_food_pb"
import { FoodType } from "@/types/Food"
import { ScriptProps } from "next/script"
import { createContext, useReducer } from "react"

type FoodsAction = 
| {
    type: "addFood"
    payload: FoodType
}
| {
    type: "updateFood"
    payload: FoodType
}
| {
    type: "removeFood"
    payload: {
        id: string
    }
}

const foodsReducer = (state: FoodType[], action: FoodsAction): FoodType[] => {
    switch (action.type) {
        case "addFood":
            return [...state, action.payload]
        case "updateFood":
            return state.map((food: FoodType) => food.id === action.payload.id ? action.payload : food)
        case "removeFood":
            return state.filter((food: FoodType) => food.id !== action.payload.id)
    }
}

function useFoodsReducer() {
    const [state, dispatch] = useReducer(foodsReducer, [
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
    ])
    return { state, dispatch }
}

const defaultFoods: ReturnType<typeof useFoodsReducer> = {
    state: [
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
    dispatch: () => {},
}

const FoodsContext = createContext(defaultFoods)
function FoodsProvider (props: ScriptProps) {
    return (
        <FoodsContext.Provider value={useFoodsReducer()}>
            {props.children}
        </FoodsContext.Provider>
    )
}

export { useFoodsReducer, defaultFoods, FoodsContext, FoodsProvider }

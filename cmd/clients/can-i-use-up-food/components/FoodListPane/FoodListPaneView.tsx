import { AmountUnit } from "@/connect/can_i_use_up_food_pb"
import FoodListItem from "@/ui-parts/FoodListItem/FoodListItem"
import { useContext, useState } from "react"

export default function FoodListPaneView() {
    // TODO: only for debugging ->
    const [foods, setFoods] = useState([
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
    // <- only for debugging
    return (
        <div className="h-full shadow-md bg-green-50">
            <div>foodlist-pane-view</div>
            {foods.map((food) =>
                <FoodListItem
                    key={food.id}
                    name={food.name}
                ></FoodListItem>
            )}
        </div>
    )
}

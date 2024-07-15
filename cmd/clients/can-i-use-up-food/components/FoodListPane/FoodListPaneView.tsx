import { AmountUnit } from "@/connect/can_i_use_up_food_pb"
import { FoodType } from "@/types/Food"
import FoodList from "@/ui-parts/FoodList/FoodList"
import { useContext, useState } from "react"

type Props = {
    foods: FoodType[]
}

export default function FoodListPaneView(props: Props) {
    const foods = props.foods
    return (
        <div className="h-full shadow-md bg-green-50">
            <div>foodlist-pane-view</div>
            <FoodList foods={foods}></FoodList>
        </div>
    )
}

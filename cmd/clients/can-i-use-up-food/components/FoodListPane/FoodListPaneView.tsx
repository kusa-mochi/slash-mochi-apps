import FoodListItem from "@/ui-parts/FoodListItem/FoodListItem"
import { useContext, useState } from "react"

export default function FoodListPaneView() {
    // TODO: only for debugging ->
    const [foods, setFoods] = useState([])
    // <- only for debugging
    return (
        <div className="h-full shadow-md bg-green-50">
            <div>foodlist-pane-view</div>
            {foods.map(() =>
                <FoodListItem></FoodListItem>
            )}
        </div>
    )
}

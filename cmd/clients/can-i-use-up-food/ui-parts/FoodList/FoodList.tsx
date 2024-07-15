import { FoodType } from "@/types/Food";
import FoodListView from "./FoodListView";

type Props = {
    foods: FoodType[]
}

export default function FoodList(props: Props) {
    return (
        <>
            <FoodListView
                foods={props.foods}
            ></FoodListView>
        </>
    )
}

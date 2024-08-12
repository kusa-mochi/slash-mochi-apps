import { FoodType } from "@/types/Food"
import FoodListItem from "../FoodListItem/FoodListItem"

type Props = {
    foods: FoodType[]
}

export default function FoodListView(props: Props) {
    return (
        <>
            {props.foods?.map((food) =>
                <div key={food.id} className="mb-2">
                    <FoodListItem name={food.name} amount={food.totalAmount} unit={food.amountUnit} />
                </div>
            )}
        </>
    )
}

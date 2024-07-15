import { FoodType } from "@/types/Food"

type Props = {
    foods: FoodType[]
}

export default function FoodListView(props: Props) {
    return (
        <div>
            {props.foods?.map((food) =>
                <div key={food.id}>
                    <input type="checkbox" id={food.id} name="food_list" /><label>{food.name}</label>
                </div>
            )}
        </div>
    )
}

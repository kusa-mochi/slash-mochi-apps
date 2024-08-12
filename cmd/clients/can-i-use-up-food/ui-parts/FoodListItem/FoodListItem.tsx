import { AmountUnit } from "@/connect/can_i_use_up_food_pb";
import FoodListItemView from "./FoodListItemView";

type Props = {
    name: string
    amount: number
    unit: AmountUnit
}

export default function FoodListItem(props: Props) {
    return (
        <>
            <FoodListItemView name={props.name} amount={props.amount} unit={props.unit} />
        </>
    )
}

import { AmountUnit } from "@/connect/can_i_use_up_food_pb";
import CalendarItemView from "./CalendarItemView";

type Props = {
    foodName: string
    amount: number
    unit: AmountUnit
}

export default function CalendarItem(props: Props) {
    return (
        <>
            <CalendarItemView foodName={props.foodName} amount={props.amount} unit={props.unit} />
        </>
    )
}

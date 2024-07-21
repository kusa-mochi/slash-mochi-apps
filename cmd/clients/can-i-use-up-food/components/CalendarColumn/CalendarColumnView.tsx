import { CalendarItemType } from "@/types/CalendarItem"
import { FoodType } from "@/types/Food"

type Props = {
    date: Date
    calendarItems: CalendarItemType[]
    foods: FoodType[]
}

export default function CalendarColumnView(props: Props) {
    console.log(`calendar items:`)
    console.log(props.calendarItems)
    return (
        <>
            <div>
                <div className="m-8">{`${props.date.getFullYear()}/${props.date.getMonth() + 1}/${props.date.getDate()}`}</div>
                <div>
                    {props.calendarItems.map((item) =>
                        <div>{props.foods.find((food) => food.id === item.foodId)?.name} {item.amount}</div>
                    )}
                </div>
            </div>
        </>
    )
}

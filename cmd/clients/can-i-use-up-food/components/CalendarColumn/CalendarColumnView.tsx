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
            <div className="flex flex-col flex-nowrap justify-start items-stretch">
                {/* date */}
                <div className="m-8 h-12 content-center">{`${props.date.getFullYear()}/${props.date.getMonth() + 1}/${props.date.getDate()}`}</div>
                {/* calendar items */}
                <div>
                    {props.calendarItems.map((item) =>
                        <div>{props.foods.find((food) => food.id === item.foodId)?.name} {item.amount}</div>
                    )}
                </div>
            </div>
        </>
    )
}

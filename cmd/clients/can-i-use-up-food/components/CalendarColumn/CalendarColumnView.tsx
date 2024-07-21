import { CalendarItemType } from "@/types/CalendarItem"

type Props = {
    date: Date
    calendarItems: CalendarItemType[]
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
                        <div>{item.foodId}</div>
                    )}
                </div>
            </div>
        </>
    )
}

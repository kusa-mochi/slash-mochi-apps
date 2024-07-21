import { CalendarItemType } from "@/types/CalendarItem"
import CalendarColumnView from "./CalendarColumnView"

type Props = {
    key: string
    date: Date
    calendarItems: CalendarItemType[]
}

export default function CalendarColumn(props: Props) {
    return (
        <>
            <CalendarColumnView date={props.date} calendarItems={props.calendarItems}></CalendarColumnView>
        </>
    )
}

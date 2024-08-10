import { CalendarItemType } from "@/types/CalendarItem"
import CalendarColumnView from "./CalendarColumnView"
import { useContext } from "react"
import { FoodsContext } from "@/context/FoodsReducer"

type Props = {
    key: string
    date: Date
    morningCalendarItems: CalendarItemType[]
    lunchCalendarItems: CalendarItemType[]
    dinnerCalendarItems: CalendarItemType[]
    otherCalendarItems: CalendarItemType[]
}

export default function CalendarColumn(props: Props) {
    const { state: foods } = useContext(FoodsContext)
    return (
        <>
            <CalendarColumnView
                date={props.date}
                morningCalendarItems={props.morningCalendarItems}
                lunchCalendarItems={props.lunchCalendarItems}
                dinnerCalendarItems={props.dinnerCalendarItems}
                otherCalendarItems={props.otherCalendarItems}
                foods={foods} />
        </>
    )
}

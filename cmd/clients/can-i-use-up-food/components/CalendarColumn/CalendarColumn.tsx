import { CalendarItemType } from "@/types/CalendarItem"
import CalendarColumnView from "./CalendarColumnView"
import { useContext } from "react"
import { FoodsContext } from "@/context/FoodsReducer"

type Props = {
    key: string
    date: Date
    calendarItems: CalendarItemType[]
}

export default function CalendarColumn(props: Props) {
    const { state: foods } = useContext(FoodsContext)
    return (
        <>
            <CalendarColumnView date={props.date} calendarItems={props.calendarItems} foods={foods}></CalendarColumnView>
        </>
    )
}

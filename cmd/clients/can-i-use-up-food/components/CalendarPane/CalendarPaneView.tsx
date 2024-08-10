import { CalendarItemType } from "@/types/CalendarItem"
import { ProjectSettingType } from "@/types/ProjectSetting"
import CalendarColumn from "../CalendarColumn/CalendarColumn"
import { MealTime } from "@/connect/can_i_use_up_food_pb"

type Props = {
    projectSettings: ProjectSettingType
    calendarItems: CalendarItemType[]
}

export default function CalendarPaneView(props: Props) {
    const startDate: Date = new Date(
        props.projectSettings.startDate.year,
        props.projectSettings.startDate.month - 1,
        props.projectSettings.startDate.date,
    )
    const endDate: Date = new Date(
        props.projectSettings.endDate.year,
        props.projectSettings.endDate.month - 1,
        props.projectSettings.endDate.date,
    )

    const calendarColumns: JSX.Element[] = []
    for (let currentDate: Date = startDate; currentDate.getTime() < endDate.getTime(); currentDate.setDate(currentDate.getDate() + 1)) {
        const date: Date = new Date(currentDate)
        const morningCalendarItems: CalendarItemType[] = props.calendarItems.filter(
            (item) => item.date.year === date.getFullYear()
            && item.date.month === date.getMonth() + 1
            && item.date.date === date.getDate()
            && item.mealTime === MealTime.Morning
        )
        const lunchCalendarItems: CalendarItemType[] = props.calendarItems.filter(
            (item) => item.date.year === date.getFullYear()
            && item.date.month === date.getMonth() + 1
            && item.date.date === date.getDate()
            && item.mealTime === MealTime.Lunch
        )
        const dinnerCalendarItems: CalendarItemType[] = props.calendarItems.filter(
            (item) => item.date.year === date.getFullYear()
            && item.date.month === date.getMonth() + 1
            && item.date.date === date.getDate()
            && item.mealTime === MealTime.Dinner
        )
        const otherCalendarItems: CalendarItemType[] = props.calendarItems.filter(
            (item) => item.date.year === date.getFullYear()
            && item.date.month === date.getMonth() + 1
            && item.date.date === date.getDate()
            && item.mealTime === MealTime.Others
        )
        calendarColumns.push(<CalendarColumn
            key={date.getTime().toString()}
            date={date}
            morningCalendarItems={morningCalendarItems}
            lunchCalendarItems={lunchCalendarItems}
            dinnerCalendarItems={dinnerCalendarItems}
            otherCalendarItems={otherCalendarItems} />)
    }

    return (
        <div className="flex flex-row flex-nowrap justify-start items-stretch h-full">
            <div className="grow-0 w-32 flex flex-col flex-nowrap justify-start items-start">
                <div className="h-28 w-px"></div>
                <div className="grow flex flex-col flex-nowrap justify-around items-center">
                    <div>朝</div>
                    <div>昼</div>
                    <div>夜</div>
                    <div>その他</div>
                </div>
            </div>
            <div className="grow overflow-x-auto scroll-smooth flex flex-row flex-nowrap justify-start items-start">
                {calendarColumns}
            </div>
        </div>
    )
}

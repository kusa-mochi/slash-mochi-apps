import { CalendarItemType } from "@/types/CalendarItem"
import { ProjectSettingType } from "@/types/ProjectSetting"
import CalendarColumn from "../CalendarColumn/CalendarColumn"

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
        const filteredCalendarItems: CalendarItemType[] = props.calendarItems.filter(
            (item) => item.date.year === date.getFullYear()
            && item.date.month === date.getMonth() + 1
            && item.date.date === date.getDate()
        )
        calendarColumns.push(<CalendarColumn key={date.getTime().toString()} date={date} calendarItems={filteredCalendarItems}></CalendarColumn>)
    }

    return (
        <>
            <div className="flex flex-row flex-nowrap justify-start items-start">
                <div className="flex flex-col flex-nowrap justify-around items-center">
                    <div>朝</div>
                    <div>昼</div>
                    <div>夜</div>
                    <div>その他</div>
                </div>
                <div className="scroll-smooth flex flex-row flex-nowrap justify-start items-center">
                    {calendarColumns}
                </div>
            </div>
        </>
    )
}

import { useContext } from "react";
import CalendarPaneView from "./CalendarPaneView";
import { CalendarItemsContext } from "@/context/CalendarItemsReducer";
import { ProjectSettingContext } from "@/context/ProjectSettingReducer";

export default function CalendarPane() {

    const { state: projectSettings } = useContext(ProjectSettingContext)
    const { state: calendarItems } = useContext(CalendarItemsContext)

    return (
        <>
            <CalendarPaneView projectSettings={projectSettings} calendarItems={calendarItems}></CalendarPaneView>
        </>
    )
}

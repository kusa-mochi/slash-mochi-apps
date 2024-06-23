import { ReactNode, useContext } from "react";
import TitleBarView from "./TitleBarView";
import { ProjectSettingContext } from "@/context/ProjectSettingReducer";

export default function TitleBar() {
    const { state: projectSettings } = useContext(ProjectSettingContext)
    return (
        <TitleBarView title={projectSettings.projectName}></TitleBarView>
    )
}

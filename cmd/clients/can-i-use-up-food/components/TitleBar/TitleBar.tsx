import { ReactNode, useContext } from "react";
import TitleBarView from "./TitleBarView";
import { Props } from "@/types/Props";
import { ProjectSettingContext } from "@/context/ProjectSettingReducer";

export default function TitleBar() {
    const { state: projectSettings } = useContext(ProjectSettingContext)
    return (
        <TitleBarView title={projectSettings.projectName}></TitleBarView>
    )
}

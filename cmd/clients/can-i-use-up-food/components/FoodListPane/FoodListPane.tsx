import { useContext } from "react";
import FoodListPaneView from "./FoodListPaneView";
import { ProjectSettingContext } from "@/context/ProjectSettingReducer";

export default function FoodListPane() {
    const { state: projectSettings } = useContext(ProjectSettingContext)
    return (
        <>
            <div>{projectSettings.projectName}</div>
            <FoodListPaneView></FoodListPaneView>
        </>
    )
}

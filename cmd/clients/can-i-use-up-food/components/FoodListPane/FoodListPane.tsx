import { useContext } from "react";
import FoodListPaneView from "./FoodListPaneView";
import { ProjectSettingContext } from "@/context/ProjectSettingReducer";
import { FoodsContext } from "@/context/FoodsReducer";

export default function FoodListPane() {
    const { state: projectSettings } = useContext(ProjectSettingContext)
    const { state: foods } = useContext(FoodsContext)
    return (
        <>
            <div>{projectSettings.projectName}</div>
            <FoodListPaneView foods={foods}></FoodListPaneView>
        </>
    )
}

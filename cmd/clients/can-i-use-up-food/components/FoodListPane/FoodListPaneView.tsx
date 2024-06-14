import { ProjectContext } from "@/context/ProjectContext"
import { useContext } from "react"

export default function FoodListPaneView() {
    const ctx = useContext(ProjectContext)
    return (
        <div className="h-full shadow-md bg-green-50">
            <div>foodlist-pane-view</div>
            <div>{ctx}</div>
        </div>
    )
}

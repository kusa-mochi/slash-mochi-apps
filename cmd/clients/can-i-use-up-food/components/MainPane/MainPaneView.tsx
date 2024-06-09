import CalendarPane from "../CalendarPane/CalendarPane";
import FoodListPane from "../FoodListPane/FoodListPane";

export default function MainPaneView() {
    return (
        <>
            <div>main-pane-view</div>
            <div>
                <FoodListPane></FoodListPane>
                <CalendarPane></CalendarPane>
            </div>
        </>
    )
}

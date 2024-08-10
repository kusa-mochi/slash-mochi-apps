import CalendarPane from "../CalendarPane/CalendarPane";
import FoodListPane from "../FoodListPane/FoodListPane";

export default function MainPaneView() {
    return (
        <div className="grow overflow-hidden flex flex-row items-stretch justify-start w-full">
            <div className="grow-0 basis-96">
                <FoodListPane></FoodListPane>
            </div>
            <div className="grow overflow-hidden">
                <CalendarPane></CalendarPane>
            </div>
        </div>
    )
}

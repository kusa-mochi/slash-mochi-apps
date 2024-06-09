import CalendarPane from "../CalendarPane/CalendarPane";
import FoodListPane from "../FoodListPane/FoodListPane";

export default function MainPaneView() {
    return (
        <div className="grow flex flex-row items-stretch justify-start w-full">
            <div className="basis-80">
                <FoodListPane></FoodListPane>
            </div>
            <div className="grow">
                <CalendarPane></CalendarPane>
            </div>
        </div>
    )
}

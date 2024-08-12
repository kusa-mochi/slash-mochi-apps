import { AmountUnit, MealTime } from "@/connect/can_i_use_up_food_pb"
import { CalendarItemType } from "@/types/CalendarItem"
import { FoodType } from "@/types/Food"
import CalendarItem from "@/ui-parts/CalendarItem/CalendarItem"

type Props = {
    date: Date
    morningCalendarItems: CalendarItemType[]
    lunchCalendarItems: CalendarItemType[]
    dinnerCalendarItems: CalendarItemType[]
    otherCalendarItems: CalendarItemType[]
    foods: FoodType[]
}

export default function CalendarColumnView(props: Props) {
    return (
        <>
            <div className="grid grid-rows-[80px_1fr_1fr_1fr_1fr] grid-cols-[200px] h-full border-r-2">
                {/* date */}
                <div className="content-center text-center">{`${props.date.getFullYear()}/${props.date.getMonth() + 1}/${props.date.getDate()}`}</div>

                {/* calendar items */}
                <div className="border-t-4 p-2">
                    {props.morningCalendarItems.map((item: CalendarItemType) => {
                        const foodItem: FoodType | undefined = props.foods.find((food: FoodType) => food.id === item.foodId)
                        if (foodItem === undefined) return null
                        return (
                            <div key={item.id} className="mb-2">
                                <CalendarItem foodName={foodItem.name} amount={item.amount} unit={foodItem.amountUnit} />
                            </div>
                        )
                    })}
                </div>
                
                <div className="border-t-2 border-dashed p-2">
                    {props.lunchCalendarItems.map((item: CalendarItemType) => {
                        const foodItem: FoodType | undefined = props.foods.find((food: FoodType) => food.id === item.foodId)
                        if (foodItem === undefined) return null
                        return (
                            <div key={item.id} className="mb-2">
                                <CalendarItem foodName={foodItem.name} amount={item.amount} unit={foodItem.amountUnit} />
                            </div>
                        )
                    })}
                </div>

                <div className="border-t-2 border-dashed p-2">
                    {props.dinnerCalendarItems.map((item: CalendarItemType) => {
                        const foodItem: FoodType | undefined = props.foods.find((food: FoodType) => food.id === item.foodId)
                        if (foodItem === undefined) return null
                        return (
                            <div key={item.id} className="mb-2">
                                <CalendarItem foodName={foodItem.name} amount={item.amount} unit={foodItem.amountUnit} />
                            </div>
                        )
                    })}
                </div>
                
                <div className="border-t-2 border-dashed p-2">
                    {props.otherCalendarItems.map((item: CalendarItemType) => {
                        const foodItem: FoodType | undefined = props.foods.find((food: FoodType) => food.id === item.foodId)
                        if (foodItem === undefined) return null
                        return (
                            <div key={item.id} className="mb-2">
                                <CalendarItem foodName={foodItem.name} amount={item.amount} unit={foodItem.amountUnit} />
                            </div>
                        )
                    })}
                </div>
            </div>
        </>
    )
}

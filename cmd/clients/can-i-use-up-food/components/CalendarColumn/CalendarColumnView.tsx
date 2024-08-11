import { MealTime } from "@/connect/can_i_use_up_food_pb"
import { CalendarItemType } from "@/types/CalendarItem"
import { FoodType } from "@/types/Food"

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
            <div className="flex flex-col flex-nowrap justify-start items-stretch h-full">
                {/* date */}
                <div className="m-8 h-12 content-center">{`${props.date.getFullYear()}/${props.date.getMonth() + 1}/${props.date.getDate()}`}</div>
                {/* calendar items */}
                <div className="grow flex flex-col flex-nowrap justify-around items-center">
                    <div>
                        {props.morningCalendarItems.map((item: CalendarItemType) => {
                            console.log(`mealtime: ${item.mealTime}, food id: ${item.foodId}`)

                            const morningFoodItem: FoodType | undefined = props.foods.find((food: FoodType) => food.id === item.foodId)
                            if (morningFoodItem === undefined) return null

                            return <div>{morningFoodItem.name} {item.amount}</div>
                        })}
                    </div>
                    
                    <div>
                        {props.lunchCalendarItems.map((item: CalendarItemType) => {
                            const morningFoodItem: FoodType | undefined = props.foods.find((food: FoodType) => food.id === item.foodId)
                            if (morningFoodItem === undefined) return null

                            return <div>{morningFoodItem.name} {item.amount}</div>
                        })}
                    </div>

                    <div>
                        {props.dinnerCalendarItems.map((item: CalendarItemType) => {
                            const morningFoodItem: FoodType | undefined = props.foods.find((food: FoodType) => food.id === item.foodId)
                            if (morningFoodItem === undefined) return null

                            return <div>{morningFoodItem.name} {item.amount}</div>
                        })}
                    </div>
                    
                    <div>
                        {props.otherCalendarItems.map((item: CalendarItemType) => {
                            const morningFoodItem: FoodType | undefined = props.foods.find((food: FoodType) => food.id === item.foodId)
                            if (morningFoodItem === undefined) return null

                            return <div>{morningFoodItem.name} {item.amount}</div>
                        })}
                    </div>
                </div>
            </div>
        </>
    )
}

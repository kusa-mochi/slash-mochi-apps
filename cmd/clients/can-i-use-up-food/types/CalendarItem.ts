import { MealTime } from "@/connect/can_i_use_up_food_pb"
import { DateType } from "./Date"

export type CalendarItemType = {
    id: string
    foodId: string
    date: DateType
    mealTime: MealTime
    amount: number
}

import { AmountUnit } from "@/connect/can_i_use_up_food_pb"
import { DateType } from "./Date"

export type FoodType = {
    id: string
    name: string
    totalAmount: number
    amountUnit: AmountUnit
    limitDate: DateType
}

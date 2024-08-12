import { AmountUnit } from "@/connect/can_i_use_up_food_pb"
import TypeConverter from "@/logical-parts/TypeConverter"

type Props = {
    name: string
    amount: number
    unit: AmountUnit
}

export default function FoodListItemView(props: Props) {
    return (
        <div className="w-full shadow-[0px_1px_2px_0px_rgba(0,0,0,0.45)] bg-green-100 rounded-full flex flex-row flex-nowrap justify-between items-center">
            <div className="grow my-2 ml-4 grid grid-cols-[1fr_40px]">
                <div className="content-center">{props.name}</div>
                <div className="text-sm text-gray-500 text-right content-center">{props.amount}&nbsp;{TypeConverter.AmountUnit2String(props.unit)}</div>
            </div>
            <div className="mx-2">︙</div>
        </div>
    )
}

import { AmountUnit } from "@/connect/can_i_use_up_food_pb"

type Props = {
    foodName: string
    amount: number
    unit: AmountUnit
}

export default function CalendarItemView(props: Props) {
    const getAmountUnitString = (unit: AmountUnit): string => {
        switch (unit) {
            case AmountUnit.NONE:
                return ""
            case AmountUnit.PIECES:
                return "個"
            case AmountUnit.MILLI_LITER:
                return "ml"
            case AmountUnit.LITER:
                return "L"
            case AmountUnit.CC:
                return "cc"
            case AmountUnit.GRAMS:
                return "g"
            case AmountUnit.KILO_GRAMS:
                return "kg"
        }
    }

    return (
        <div className="w-full shadow-[0px_1px_2px_0px_rgba(0,0,0,0.45)] bg-green-100 rounded-full flex flex-row flex-nowrap justify-between items-center">
            <div className="grow my-px ml-4 grid grid-cols-[1fr_40px]">
                <div className="content-center">{props.foodName}</div>
                <div className="text-sm text-gray-500 text-right content-center">{props.amount}&nbsp;{getAmountUnitString(props.unit)}</div>
            </div>
            <div className="mx-2">︙</div>
        </div>
    )
}

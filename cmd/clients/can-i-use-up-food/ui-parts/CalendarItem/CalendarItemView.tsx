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
        <div className="w-full shadow-md bg-green-200 rounded-full flex flex-row flex-nowrap justify-between items-center">
            <div className="grow">{props.foodName} {props.amount} {getAmountUnitString(props.unit)}</div>
            <div>︙</div>
        </div>
    )
}

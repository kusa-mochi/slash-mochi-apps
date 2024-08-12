import { AmountUnit } from "@/connect/can_i_use_up_food_pb";

export default abstract class TypeConverter {
    public static AmountUnit2String(unit: AmountUnit): string {
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
}

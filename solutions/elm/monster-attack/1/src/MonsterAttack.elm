module MonsterAttack exposing (..)


type alias MonsterDamage =
    String

attackFormat : MonsterDamage -> MonsterDamage -> Int -> MonsterDamage
attackFormat format monsterDamage strength =
    "."
    |> String.append (String.fromInt strength)
    |> String.append format
    |> String.append monsterDamage


type Weapon = Sword | Claw


weaponFormat : Weapon -> String
weaponFormat weapon =
    let
        weaponString : String
        weaponString = case weapon of
            Sword -> "sword"
            Claw -> "claw"
    in
        "Attacked with " ++ weaponString ++ " of strength "


attackWithSword1 : MonsterDamage -> Int -> MonsterDamage
attackWithSword1 = attackFormat (weaponFormat Sword)


attackWithClaw1 : MonsterDamage -> Int -> MonsterDamage
attackWithClaw1 = attackFormat (weaponFormat Claw)


attack1 : MonsterDamage -> MonsterDamage
attack1 =
    let
        annalynAttack : MonsterDamage -> MonsterDamage
        annalynAttack dmg = attackWithSword1 dmg 5
        kazakAttack : MonsterDamage -> MonsterDamage
        kazakAttack dmg = attackWithClaw1 dmg 1
    in
        annalynAttack >> kazakAttack >> kazakAttack >> annalynAttack


newAttackFormat : MonsterDamage -> Int -> MonsterDamage -> MonsterDamage
newAttackFormat format strength monsterDamage = attackFormat format monsterDamage strength


attackWithSword2 : Int -> MonsterDamage -> MonsterDamage
attackWithSword2 = newAttackFormat (weaponFormat Sword)


attackWithClaw2 : Int -> MonsterDamage -> MonsterDamage
attackWithClaw2 = newAttackFormat (weaponFormat Claw)


attack2 : MonsterDamage -> MonsterDamage
attack2 =
    let
        annalynAttack : MonsterDamage -> MonsterDamage
        annalynAttack = attackWithSword2 5
        kazakAttack : MonsterDamage -> MonsterDamage
        kazakAttack = attackWithClaw2 1
    in
        annalynAttack >> kazakAttack >> kazakAttack >> annalynAttack


attack3 : MonsterDamage -> MonsterDamage
attack3 = attack2

package components

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	list "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const translationsRaw = `
{
  "categories": [
    {
      "id": "weapons",
      "items": {
        "Weapon_Sword_001": "Traveler's Sword",
        "Weapon_Sword_002": "Soldier's Broadsword",
        "Weapon_Sword_003": "Knight's Broadsword",
        "Weapon_Sword_004": "Boko Club",
        "Weapon_Sword_005": "Spiked Boko Club",
        "Weapon_Sword_006": "Dragonbone Boko Club",
        "Weapon_Sword_007": "Lizal Boomerand",
        "Weapon_Sword_008": "Lizal Forked Boomerang",
        "Weapon_Sword_009": "Lizal Tri-Boomerang",
        "Weapon_Sword_013": "Guardian Sword",
        "Weapon_Sword_014": "Guardian Sword+",
        "Weapon_Sword_015": "Guardian Sword++",
        "Weapon_Sword_016": "Lynel Sword",
        "Weapon_Sword_017": "Mighty Lynel Sword",
        "Weapon_Sword_018": "Savage Lynel Sword",
        "Weapon_Sword_019": "Bokoblin Arm",
        "Weapon_Sword_020": "Lizalfos Arm",
        "Weapon_Sword_021": "Rusty Broadsword",
        "Weapon_Sword_022": "Soup ladle",
        "Weapon_Sword_023": "Ancient Short Sword",
        "Weapon_Sword_024": "Royal Broadsword",
        "Weapon_Sword_025": "Forest Dweller's Sword",
        "Weapon_Sword_027": "Zora Sword",
        "Weapon_Sword_029": "Gerudo Scimitar",
        "Weapon_Sword_030": "Moonlight Scimitar",
        "Weapon_Sword_031": "Feathered Edge",
        "Weapon_Sword_033": "Flameblade",
        "Weapon_Sword_034": "Frostblade",
        "Weapon_Sword_035": "Thunderblade",
        "Weapon_Sword_040": "Spring-Loaded Hammer",
        "Weapon_Sword_041": "Eightfold Blade",
        "Weapon_Sword_043": "Torch",
        "Weapon_Sword_044": "Tree Branch",
        "Weapon_Sword_047": "Royal Guard's Sword",
        "Weapon_Sword_048": "Meteor Rod",
        "Weapon_Sword_049": "Blizzard Rod",
        "Weapon_Sword_050": "Thunderstorm Rod",
        "Weapon_Sword_051": "Boomerang",
        "Weapon_Sword_052": "Scimitar of the Seven",
        "Weapon_Sword_053": "Vicious Sickle",
        "Weapon_Sword_056": "Master Sword (Broken/Unequippable)",
        "Weapon_Sword_057": "Goddess Sword",
        "Weapon_Sword_058": "Hero's Sword (amiibo 8-bit)",
        "Weapon_Sword_059": "Sea-Breeze Boomerang (amiibo WW)",
        "Weapon_Sword_060": "Fire Rod",
        "Weapon_Sword_061": "Ice Rod",
        "Weapon_Sword_062": "Lightning Rod",
        "Weapon_Sword_070": "Master Sword",
        "Weapon_Sword_071": "Master Sword (no near malice, no charge)",
        "Weapon_Sword_072": "Master Sword (near malice, no charge)",
        "Weapon_Sword_073": "Demon Carver",
        "Weapon_Sword_500": "Lantern",
        "Weapon_Lsword_001": "Traveler's Claymore",
        "Weapon_Lsword_002": "Soldier's Claymore",
        "Weapon_Lsword_003": "Knight's Claymore",
        "Weapon_Lsword_004": "Boko Bat",
        "Weapon_Lsword_005": "Spiked Boko Bat",
        "Weapon_Lsword_006": "Dragonbone Boko Bat",
        "Weapon_Lsword_010": "Moblin Club",
        "Weapon_Lsword_011": "Spiked Moblin Club",
        "Weapon_Lsword_012": "Dragonbone Moblin Club",
        "Weapon_Lsword_013": "Ancient Battle Axe",
        "Weapon_Lsword_014": "Ancient Battle Axe+",
        "Weapon_Lsword_015": "Ancient Battle Axe++",
        "Weapon_Lsword_016": "Lynel Crusher",
        "Weapon_Lsword_017": "Mighty Lynel Crusher",
        "Weapon_Lsword_018": "Savage Lynel Crusher",
        "Weapon_Lsword_019": "Moblin Arm",
        "Weapon_Lsword_020": "Rusty Claymore",
        "Weapon_Lsword_023": "Ancient Bladesaw",
        "Weapon_Lsword_024": "Royal Claymore",
        "Weapon_Lsword_027": "Silver Longsword",
        "Weapon_Lsword_029": "Golden Claymore",
        "Weapon_Lsword_030": "Double Axe",
        "Weapon_Lsword_031": "Iron Sledgehammer",
        "Weapon_Lsword_032": "Woodcutter's Axe",
        "Weapon_Lsword_033": "Great Flameblade",
        "Weapon_Lsword_034": "Great Frostblade",
        "Weapon_Lsword_035": "Great Thunderblade",
        "Weapon_Lsword_036": "Cobble Crusher",
        "Weapon_Lsword_037": "Stone Smasher",
        "Weapon_Lsword_038": "Boat Oar",
        "Weapon_Lsword_041": "Eightfold Longblade",
        "Weapon_Lsword_045": "Farming Hoe",
        "Weapon_Lsword_047": "Royal Guard's Claymore",
        "Weapon_Lsword_051": "Giant Boomerang",
        "Weapon_Lsword_054": "Boulder Breaker",
        "Weapon_Lsword_055": "Edge of Duality",
        "Weapon_Lsword_056": "Korok Leaf",
        "Weapon_Lsword_057": "Sword of the Six Sages (amiibo TP)",
        "Weapon_Lsword_059": "Biggoron's Sword (amiibo OoT)",
        "Weapon_Lsword_060": "Fierce Deity Sword (amiibo MM)",
        "Weapon_Lsword_074": "Windcleaver",
        "Weapon_Spear_001": "Traveler's Spear",
        "Weapon_Spear_002": "Soldier's Spear",
        "Weapon_Spear_003": "Knight's Halberd",
        "Weapon_Spear_004": "Boko Spear",
        "Weapon_Spear_005": "Spiked Boko Spear",
        "Weapon_Spear_006": "Dragonbone Boko Spear",
        "Weapon_Spear_007": "Lizal Spear",
        "Weapon_Spear_008": "Enhanced Lizal Spear",
        "Weapon_Spear_009": "Forked Lizal Spear",
        "Weapon_Spear_010": "Moblin Spear",
        "Weapon_Spear_011": "Spiked Moblin Spear",
        "Weapon_Spear_012": "Dragonbone Moblin Spear",
        "Weapon_Spear_013": "Guardian Spear",
        "Weapon_Spear_014": "Guardian Spear+",
        "Weapon_Spear_015": "Guardian Spear++",
        "Weapon_Spear_016": "Lynel Spear",
        "Weapon_Spear_017": "Mighty Lynel Spear",
        "Weapon_Spear_018": "Savage Lynel Spear",
        "Weapon_Spear_021": "Rusty Halberd",
        "Weapon_Spear_022": "Farmer's Pichfork",
        "Weapon_Spear_023": "Ancient Spear",
        "Weapon_Spear_024": "Royal Halberd",
        "Weapon_Spear_025": "Forest Dweller's Spear",
        "Weapon_Spear_027": "Zora Spear",
        "Weapon_Spear_028": "Silverscale Spear",
        "Weapon_Spear_029": "Gerudo Spear",
        "Weapon_Spear_030": "Throwing Spear",
        "Weapon_Spear_031": "Drillshaft",
        "Weapon_Spear_032": "Feathered Spear",
        "Weapon_Spear_033": "Flamespear",
        "Weapon_Spear_034": "Frostspear",
        "Weapon_Spear_035": "Thunderspear",
        "Weapon_Spear_036": "Wooden Mop",
        "Weapon_Spear_037": "Serpentine Spear",
        "Weapon_Spear_038": "Fishing Harpoon",
        "Weapon_Spear_047": "Royal Guard's Spear",
        "Weapon_Spear_049": "Ceremonial Trident",
        "Weapon_Spear_050": "Lightscale Trident",
        "Weapon_Sword_502": "One-Hit Obliterator (v1.4)"
      }
    },
    {
      "id": "bows",
      "items": {
        "Weapon_Bow_001": "Traveler's Bow",
        "Weapon_Bow_002": "Soldier's Bow",
        "Weapon_Bow_003": "Spiked Boko Bow",
        "Weapon_Bow_004": "Boko Bow",
        "Weapon_Bow_006": "Lizal Bow",
        "Weapon_Bow_009": "Lynel Bow",
        "Weapon_Bow_011": "Strengthened Lizal Bow",
        "Weapon_Bow_013": "Forest Dweller's Bow",
        "Weapon_Bow_014": "Silver Bow",
        "Weapon_Bow_015": "Golden Bow",
        "Weapon_Bow_016": "Swallow Bow",
        "Weapon_Bow_017": "Falcon Bow",
        "Weapon_Bow_023": "Ancient Bow",
        "Weapon_Bow_026": "Mighty Lynel Bow",
        "Weapon_Bow_027": "Dragon Bone Boko Bow",
        "Weapon_Bow_028": "Great Eagle Bow",
        "Weapon_Bow_029": "Phrenic Bow",
        "Weapon_Bow_030": "Steel Lizal Bow",
        "Weapon_Bow_032": "Savage Lynel Bow",
        "Weapon_Bow_033": "Royal Guard's Bow",
        "Weapon_Bow_035": "Knight's Bow",
        "Weapon_Bow_036": "Royal Bow",
        "Weapon_Bow_038": "Wooden Bow",
        "Weapon_Bow_040": "Duplex Bow",
        "Weapon_Bow_071": "Bow of Light",
        "Weapon_Bow_072": "Twilight Bow (amiibo TP)",
        "NormalArrow": "Arrow",
        "FireArrow": "Fire Arrow",
        "IceArrow": "Ice Arrow",
        "ElectricArrow": "Shock Arrow",
        "BombArrow_A": "Bomb Arrow",
        "AncientArrow": "Ancient Arrow"
      }
    },
    {
      "id": "shields",
      "items": {
        "Weapon_Shield_001": "Wooden Shield",
        "Weapon_Shield_002": "Soldier's Shield",
        "Weapon_Shield_003": "Knight's Shield",
        "Weapon_Shield_004": "Boko Shield",
        "Weapon_Shield_005": "Spiked Boko Shield",
        "Weapon_Shield_006": "Dragonbone Boko Shield",
        "Weapon_Shield_007": "Lizal Shield",
        "Weapon_Shield_008": "Reinforced Lizal Shield",
        "Weapon_Shield_009": "Steel Lizal Shield",
        "Weapon_Shield_013": "Guardian Shield",
        "Weapon_Shield_014": "Guardian Shield+",
        "Weapon_Shield_015": "Guardian Shield++",
        "Weapon_Shield_016": "Lynel Shield",
        "Weapon_Shield_017": "Mighty Lynel Shield",
        "Weapon_Shield_018": "Savage Lynel Shield",
        "Weapon_Shield_021": "Rusty Shield",
        "Weapon_Shield_022": "Royal Shield",
        "Weapon_Shield_023": "Forest Dweller's Shield",
        "Weapon_Shield_025": "Silver Shield",
        "Weapon_Shield_026": "Gerudo Shield",
        "Weapon_Shield_030": "Hylian Shield",
        "Weapon_Shield_031": "Hunter's Shield",
        "Weapon_Shield_032": "Fisherman's Shield",
        "Weapon_Shield_033": "Royal Guard's Shield",
        "Weapon_Shield_034": "Emblazoned Shield",
        "Weapon_Shield_035": "Traveler's Shield",
        "Weapon_Shield_036": "Radiant Shield",
        "Weapon_Shield_037": "Daybreaker",
        "Weapon_Shield_038": "Ancient Shield",
        "Weapon_Shield_040": "Pot Lid",
        "Weapon_Shield_041": "Shield of the Mind's Eye",
        "Weapon_Shield_042": "Kite Shield",
        "Weapon_Shield_057": "Hero's Shield (Wind Waker)"
      }
    },
    {
      "id": "materials",
      "items": {
        "Item_Fruit_A": "Apple",
        "Item_Fruit_B": "Wildberry",
        "Item_Fruit_C": "Voltfruit",
        "Item_Fruit_D": "Hearty Durian",
        "Item_Fruit_E": "Fleet-Lotus Seeds",
        "Item_Fruit_E_00": "Fleet-Lotus Seeds x0",
        "Item_Fruit_F": "Hydromelon",
        "Item_Fruit_G": "Palm Fruit",
        "Item_Fruit_H": "Mighty Bananas",
        "Item_Fruit_I": "Spicy Pepper",
        "Item_Fruit_J": "Fortified Pumpkin",
        "Item_Fruit_K": "Acorn",
        "Item_Fruit_L": "Chickaloo Tree Nut",
        "Item_Mushroom_A": "Stamella Mushroom",
        "Item_MushroomGet_A": "Stamella Mushroom",
        "Item_Mushroom_B": "Chillshroom",
        "Item_MushroomGet_B": "Chillshroom",
        "Item_Mushroom_C": "Sunshroom",
        "Item_MushroomGet_C": "Sunshroom",
        "Item_Mushroom_D": "Rushroom",
        "Item_MushroomGet_D": "Rushroom",
        "Item_Mushroom_E": "Hylian Mushroom",
        "Item_MushroomGet_E": "Hylian Mushroom",
        "Item_Mushroom_F": "Hearty Truffle",
        "Item_Mushroom_F_00": "Hearty Truffle x0",
        "Item_MushroomGet_F": "Hearty Truffle",
        "Item_Mushroom_H": "Zapshroom",
        "Item_MushroomGet_H": "Zapshroom",
        "Item_Mushroom_J": "Silent Shroom",
        "Item_MushroomGet_J": "Silent Shroom",
        "Item_Mushroom_L": "Razorshroom",
        "Item_MushroomGet_L": "Razorshroom",
        "Item_Mushroom_M": "Ironshroom",
        "Item_MushroomGet_M": "Ironshroom",
        "Item_Mushroom_N": "Big Hearty Truffle",
        "Item_MushroomGet_N": "Big Hearty Truffle",
        "Item_Mushroom_N_00": "Big Hearty Truffle x0",
        "Item_Mushroom_O": "Endura Shroom",
        "Item_MushroomGet_O": "Endura Shroom",
        "Item_Plant_A": "Hyrule Herb",
        "Item_PlantGet_A": "Hyrule Herb",
        "Item_Plant_B": "Hearty Radish",
        "Item_PlantGet_B": "Hearty Radish",
        "Item_Plant_C": "Big Hearty Radish",
        "Item_PlantGet_C": "Big Hearty Radish",
        "Item_Plant_E": "Cool Safflina",
        "Item_PlantGet_E": "Cool Safflina",
        "Item_Plant_F": "Warm Safflina",
        "Item_PlantGet_F": "Warm Safflina",
        "Item_Plant_G": "Mighty Thistle",
        "Item_PlantGet_G": "Mighty Thistle",
        "Item_Plant_H": "Armoranth",
        "Item_PlantGet_H": "Armoranth",
        "Item_Plant_I": "Blue Nightshade",
        "Item_PlantGet_I": "Blue Nightshade",
        "Item_Plant_J": "Silent Princess",
        "Item_PlantGet_J": "Silent Princess",
        "Item_Plant_L": "Electric Safflina",
        "Item_PlantGet_L": "Electric Safflina",
        "Item_Plant_M": "Swift Carrot",
        "Item_PlantGet_M": "Swift Carrot",
        "Item_Plant_O": "Swift Violet",
        "Item_PlantGet_O": "Swift Violet",
        "Item_Plant_Q": "Endura Carrot",
        "Item_PlantGet_Q": "Endura Carrot",
        "Item_Meat_01": "Raw Meat",
        "Item_Meat_02": "Raw Prime Meat",
        "Item_Meat_06": "Raw Bird Drumstick",
        "Item_Meat_07": "Raw Bird Thigh",
        "Item_Meat_11": "Raw Gourmet Meat",
        "Item_Meat_12": "Raw Whole Meat",
        "Item_FishGet_A": "Hylian Bass",
        "Item_FishGet_B": "Hearty Bass",
        "Item_FishGet_C": "Chillfin Trout",
        "Item_FishGet_D": "Voltfin Trout",
        "Item_FishGet_E": "Mighty Carp",
        "Item_FishGet_F": "Mighty Porgy",
        "Item_FishGet_G": "Armored Porgy",
        "Item_FishGet_H": "Armored Carp",
        "Item_FishGet_I": "Hearty Salmon",
        "Item_FishGet_J": "Sizzlefin Trout",
        "Item_FishGet_K": "Hearty Blueshell Snail",
        "Item_FishGet_L": "Staminoka Bass",
        "Item_FishGet_M": "Sneaky River Snail",
        "Item_FishGet_X": "Stealthfin Trout",
        "Item_FishGet_Z": "Sanke Carp",
        "Animal_Insect_A": "Hot-Footed Frog",
        "Item_InsectGet_A": "Hot-Footed Frog",
        "Animal_Insect_B": "Tireless Frog",
        "Item_InsectGet_B": "Tireless Frog",
        "Animal_Insect_C": "Cold Darner",
        "Item_InsectGet_C": "Cold Darner",
        "Animal_Insect_E": "Sunset Firefly",
        "Item_InsectGet_E": "Sunset Firefly",
        "Animal_Insect_F": "Fairy",
        "Item_InsectGet_F": "Fairy",
        "Animal_Insect_G": "Bladed Rhino Beetle",
        "Item_InsectGet_G": "Bladed Rhino Beetle",
        "Animal_Insect_H": "Restless Cricket",
        "Item_InsectGet_H": "Restless Cricket",
        "Animal_Insect_I": "Electric Darner",
        "Item_InsectGet_I": "Electric Darner",
        "Animal_Insect_K": "Razorclaw Crab",
        "Item_InsectGet_K": "Razorclaw Crab",
        "Animal_Insect_M": "Hearty Lizard",
        "Item_InsectGet_M": "Hearty Lizard",
        "Animal_Insect_N": "Winterwing Butterfly",
        "Item_InsectGet_N": "Winterwing Butterfly",
        "Animal_Insect_O": "Ironshell Crab",
        "Item_InsectGet_O": "Ironshell Crab",
        "Animal_Insect_P": "Rugged Rhino Beetle",
        "Item_InsectGet_P": "Rugged Rhino Beetle",
        "Animal_Insect_Q": "Summerwing Butterfly",
        "Item_InsectGet_Q": "Summerwing Butterfly",
        "Animal_Insect_R": "Thunderwing Butterfly",
        "Item_InsectGet_R": "Thunderwing Butterfly",
        "Animal_Insect_S": "Hightail Lizard",
        "Item_InsectGet_S": "Hightail Lizard",
        "Animal_Insect_T": "Warm Darner",
        "Item_InsectGet_T": "Warm Darner",
        "Animal_Insect_X": "Fireproof Lizard",
        "Item_InsectGet_X": "Fireproof Lizard",
        "Animal_Insect_Z": "Bright-Eyed Crab",
        "Item_InsectGet_Z": "Bright-Eyed Crab",
        "Animal_Insect_AA": "Energetic Rhino Beetle",
        "Item_InsectGet_AA": "Energetic Rhino Beetle",
        "Animal_Insect_AB": "Smotherwing Butterfly",
        "Item_InsectGet_AB": "Smotherwing Butterfly",
        "BeeHome": "Courser Bee Honey",
        "Obj_FireWoodBundle": "Wood",
        "Item_Enemy_00": "Bokoblin Horn",
        "Item_Enemy_01": "Bokoblin Fang",
        "Item_Enemy_02": "Bokoblin Guts",
        "Item_Enemy_03": "Lizalfos Horn",
        "Item_Enemy_04": "Lizalfos Talon",
        "Item_Enemy_05": "Lizalfos Tail",
        "Item_Enemy_06": "Moblin Horn",
        "Item_Enemy_07": "Moblin Fang",
        "Item_Enemy_08": "Moblin Guts",
        "Item_Enemy_12": "Lynel Horn",
        "Item_Enemy_13": "Lynel Hoof",
        "Item_Enemy_14": "Lynel Guts",
        "Item_Enemy_15": "Red Chuchu Jelly",
        "Item_Enemy_16": "Yellow Chuchu Jelly",
        "Item_Enemy_17": "White Chuchu Jelly",
        "Item_Enemy_18": "Keese Wing",
        "Item_Enemy_19": "Keese Eyeball",
        "Item_Enemy_20": "Octorok Tentacle",
        "Item_Enemy_21": "Tentacle Eyeball",
        "Item_Enemy_24": "Molduga Fin",
        "Item_Enemy_25": "Molduga Guts",
        "Item_Enemy_26": "Ancient Gear",
        "Item_Enemy_27": "Ancient Screw",
        "Item_Enemy_28": "Ancient Spring",
        "Item_Enemy_29": "Ancient Shaft",
        "Item_Enemy_30": "Ancient Core",
        "Item_Enemy_31": "Giant Ancient Core",
        "Item_Enemy_32": "Hinox Toenail",
        "Item_Enemy_33": "Hinox Tooth",
        "Item_Enemy_34": "Hinox Guts",
        "Item_Enemy_38": "Dinraal's Scale",
        "Item_Enemy_39": "Dinraal's Claw",
        "Item_Enemy_40": "Chuchu Jelly",
        "Item_Enemy_41": "Red Lizalfos Tail",
        "Item_Enemy_42": "Icy Lizalfos Tail",
        "Item_Enemy_43": "Yellow Lizalfos Tail",
        "Item_Enemy_44": "Fire Keese Wing",
        "Item_Enemy_45": "Electric Keese Wing",
        "Item_Enemy_46": "Ice Keese Wing",
        "Item_Enemy_47": "Shard of Dinraal's Fang",
        "Item_Enemy_48": "Shard of Dinraal's Horn",
        "Item_Enemy_49": "Naydra's Scale",
        "Item_Enemy_50": "Naydra's Claw",
        "Item_Enemy_51": "Shard of Naydra's Fang",
        "Item_Enemy_52": "Shard of Naydra's Horn",
        "Item_Enemy_53": "Farosh's Scale",
        "Item_Enemy_54": "Farosh's Claw",
        "Item_Enemy_55": "Shard of Farosh's Fang",
        "Item_Enemy_56": "Shard of Farosh's Horn",
        "Item_Enemy_57": "Octo Balloon",
        "Item_Enemy_Put_57": "Octo Balloon",
        "Item_Material_01": "Cane Sugar",
        "Item_Material_02": "Goron Spice",
        "Item_Material_03": "Hylian Rice",
        "Item_Material_04": "Bird Egg",
        "Item_Material_05": "Fresh Milk",
        "Item_Material_06": "Goat Butter",
        "Item_Material_07": "Tabantha Wheat",
        "Item_Material_08": "Monster Extract",
        "Item_Ore_A": "Diamond",
        "Item_Ore_A_00": "Diamant x0",
        "Item_Ore_B": "Ruby",
        "Item_Ore_C": "Sapphire",
        "Item_Ore_D": "Topaz",
        "Item_Ore_E": "Opal",
        "Item_Ore_F": "Amber",
        "Item_Ore_G": "Luminous Stone",
        "Item_Ore_H": "Rock Salt",
        "Item_Ore_I": "Flint",
        "Item_Ore_J": "Star Fragment"
      }
    },
    {
      "id": "food",
      "items": {
        "Item_Boiled_01": "Hard-Boiled Egg",
        "Item_ChilledFish_01": "Frozen Bass",
        "Item_ChilledFish_02": "Frozen Hearty Salmon",
        "Item_ChilledFish_03": "Frozen Trout",
        "Item_ChilledFish_04": "Frozen Carp",
        "Item_ChilledFish_05": "Frozen Porgy",
        "Item_ChilledFish_06": "Frozen Hearty Bass",
        "Item_ChilledFish_07": "Frozen Crab",
        "Item_ChilledFish_08": "Frozen River Snail",
        "Item_ChilledFish_09": "Icy Hearty Blueshell Snail",
        "Item_RoastFish_01": "Roasted Bass",
        "Item_RoastFish_02": "Roasted Hearty Bass",
        "Item_RoastFish_03": "Roasted Trout",
        "Item_RoastFish_04": "Roasted Hearty Salmon",
        "Item_RoastFish_07": "Roasted Carp",
        "Item_RoastFish_09": "Roasted Porgy",
        "Item_RoastFish_11": "Blueshell Escargot",
        "Item_RoastFish_13": "Sneaky River Escargot",
        "Item_RoastFish_15": "Blackened Crab",
        "Item_Roast_01": "Seared Steak",
        "Item_Roast_02": "Roasted Bird Drumstick",
        "Item_Roast_03": "Baked Apple",
        "Item_Roast_04": "Toasty Stamella Shroom",
        "Item_Roast_05": "Toasted Hearty Truffle",
        "Item_Roast_06": "Toasty Hylian Shroom",
        "Item_Roast_07": "Roasted Wildberry",
        "Item_Roast_08": "Roasted Voltfruit",
        "Item_Roast_09": "Roasted Hearty Durian",
        "Item_Roast_10": "Baked Palm Fruit",
        "Item_Roast_11": "Roasted Mighty Bananas",
        "Item_Roast_12": "Roasted Hydromelon",
        "Item_Roast_13": "Charred Pepper",
        "Item_Roast_15": "Baked Fortified Pumpkin",
        "Item_Roast_16": "Roasted Lotus Seed",
        "Item_Roast_18": "Roasted Radish",
        "Item_Roast_19": "Roasted Big Radish",
        "Item_Roast_24": "Roasted Swift Carrot",
        "Item_Roast_27": "Roasted Mighty Thistle",
        "Item_Roast_28": "Roasted Armoranth",
        "Item_Roast_31": "Toasty Chillshroom",
        "Item_Roast_32": "Toasty Sunshroom",
        "Item_Roast_33": "Toasty Zapshroom",
        "Item_Roast_36": "Toasty Rushroom",
        "Item_Roast_37": "Toasty Razorshroom",
        "Item_Roast_38": "Toasty Ironshroom",
        "Item_Roast_39": "Toasty Silent Shroom",
        "Item_Roast_40": "Seared Prime Steak",
        "Item_Roast_41": "Roasted Bird Thigh",
        "Item_Roast_45": "Seared Gourmet Steak",
        "Item_Roast_46": "Roasted Whole Bird",
        "Item_Roast_48": "Roasted Acorn",
        "Item_Roast_49": "Toasted Big Heart Truffle",
        "Item_Roast_50": "Roasted Endura Carrot",
        "Item_Roast_51": "Campfire Egg",
        "Item_Roast_52": "Roasted Tree Nut",
        "Item_Roast_53": "Toasty Endura Shroom",
        "Item_Chilled_01": "Icy Meat",
        "Item_Chilled_02": "Icy Prime Meat",
        "Item_Chilled_03": "Icy Gourmet Meat",
        "Item_Chilled_04": "Frozen Bird Drumstick",
        "Item_Chilled_05": "Frozen Bird Thigh",
        "Item_Chilled_06": "Frozen Whole Bird",
        "Item_Cook_A_01": "Mushroom Skewer",
        "Item_Cook_A_02": "Steamed Mushrooms",
        "Item_Cook_A_03": "Steamed Fruit",
        "Item_Cook_A_04": "Steamed Fish",
        "Item_Cook_A_05": "Steamed Meat",
        "Item_Cook_A_07": "Fruit and Mushroom Mix",
        "Item_Cook_A_08": "Fish and Mushroom Skewer",
        "Item_Cook_A_09": "Meat and Mushroom Skewer",
        "Item_Cook_A_10": "Omelet",
        "Item_Cook_A_11": "Glazed Mushroom",
        "Item_Cook_A_12": "Glazed Meat",
        "Item_Cook_A_13": "Glazed Seafood",
        "Item_Cook_A_14": "Glazed Veggies",
        "Item_Cook_B_01": "Fried Wild Greens",
        "Item_Cook_B_02": "Simmered Fruits",
        "Item_Cook_B_05": "Fish Skewer",
        "Item_Cook_B_06": "Meat Skewer",
        "Item_Cook_B_11": "Copious Fried Wild Greens",
        "Item_Cook_B_12": "Copious Simmered Fruits",
        "Item_Cook_B_13": "Copious Mushroom Skewers",
        "Item_Cook_B_15": "Copious Seafood Skewers",
        "Item_Cook_B_16": "Copious Meat Skewers",
        "Item_Cook_B_17": "Meat and Seafood Fry",
        "Item_Cook_B_18": "Prime Meat and Seafood Fry",
        "Item_Cook_B_19": "Gourmet Meat and Seafood Fry",
        "Item_Cook_B_20": "Meat-Stuffed Pumpkin",
        "Item_Cook_B_21": "Sautéed Peppers",
        "Item_Cook_B_22": "Sautéed Nuts",
        "Item_Cook_B_23": "Seafood Skewers",
        "Item_Cook_C_16": "Fairy Tonic",
        "Item_Cook_C_17": "Elixir",
        "Item_Cook_D_01": "Salt-Grilled Mushrooms",
        "Item_Cook_D_02": "Salt-Grilled Greens",
        "Item_Cook_D_03": "Salt-Grilled Fish",
        "Item_Cook_D_04": "Salt-Grilled Meat",
        "Item_Cook_D_05": "Salt-Grilled Prime Meat",
        "Item_Cook_D_06": "Salt-Grilled Gourmet Meat",
        "Item_Cook_D_07": "Pepper Steak",
        "Item_Cook_D_08": "Pepper Seafood",
        "Item_Cook_D_09": "Salt-Grilled Crab",
        "Item_Cook_D_10": "Crab Stir-Fry",
        "Item_Cook_E_01": "Poultry Pilaf",
        "Item_Cook_E_02": "Prime Poultry Pilaf",
        "Item_Cook_E_03": "Gourmet Poultry Pilaf",
        "Item_Cook_E_04": "Fried Egg and Rice",
        "Item_Cook_F_01": "Creamy Meat Soup",
        "Item_Cook_F_02": "Creamy Seafood Soup",
        "Item_Cook_F_03": "Veggie Cream Soup",
        "Item_Cook_F_04": "Creamy Heart Soup",
        "Item_Cook_G_02": "Seafood Rice Balls",
        "Item_Cook_G_03": "Veggie Rice Balls",
        "Item_Cook_G_04": "Mushroom Rice Balls",
        "Item_Cook_G_05": "Meat and Rice Bowl",
        "Item_Cook_G_06": "Prime Meat and Rice Bowl",
        "Item_Cook_G_09": "Gourmet Meat and Rice Bowl",
        "Item_Cook_G_10": "Seafood Fried Rice",
        "Item_Cook_G_11": "Curry Pilaf",
        "Item_Cook_G_12": "Mushroom Risotto",
        "Item_Cook_G_13": "Vegetable Risotto",
        "Item_Cook_G_14": "Salmon Risotto",
        "Item_Cook_G_15": "Meaty Rice Balls",
        "Item_Cook_G_16": "Crab Omelet with Rice",
        "Item_Cook_G_17": "Crab Risotto",
        "Item_Cook_H_01": "Seafood Meunière",
        "Item_Cook_H_02": "Porgy Meunière",
        "Item_Cook_H_03": "Salmon Meunière",
        "Item_Cook_I_01": "Fruit Pie",
        "Item_Cook_I_02": "Apple Pie",
        "Item_Cook_I_03": "Egg Tart",
        "Item_Cook_I_04": "Meat Pie",
        "Item_Cook_I_05": "Carrot Cake",
        "Item_Cook_I_06": "Pumpkin Pie",
        "Item_Cook_I_07": "Hot Buttered Apple",
        "Item_Cook_I_08": "Honeyed Apple",
        "Item_Cook_I_09": "Honeyed Fruits",
        "Item_Cook_I_10": "Plain Crepe",
        "Item_Cook_I_11": "Wildberry Crepe",
        "Item_Cook_I_12": "Nutcake",
        "Item_Cook_I_13": "Fried Bananas",
        "Item_Cook_I_14": "Egg Pudding",
        "Item_Cook_I_15": "Fish Pie",
        "Item_Cook_I_16": "Honey Candy",
        "Item_Cook_I_17": "Honey Crepe",
        "Item_Cook_J_01": "Curry Rice",
        "Item_Cook_J_02": "Vegetable Curry",
        "Item_Cook_J_03": "Seafood Curry",
        "Item_Cook_J_04": "Poultry Curry",
        "Item_Cook_J_05": "Prime Poultry Curry",
        "Item_Cook_J_06": "Meat Curry",
        "Item_Cook_J_07": "Prime Meat Curry",
        "Item_Cook_J_08": "Gourmet Poultry Curry",
        "Item_Cook_J_09": "Gourmet Meat Curry",
        "Item_Cook_K_01": "Meat Stew",
        "Item_Cook_K_02": "Prime Meat Stew",
        "Item_Cook_K_03": "Pumpkin Stew",
        "Item_Cook_K_04": "Clam Chowder",
        "Item_Cook_K_05": "Gourmet Meat Stew",
        "Item_Cook_K_06": "Cream of Mushroom Soup",
        "Item_Cook_K_07": "Cream of Vegetable Soup",
        "Item_Cook_K_08": "Carrot Stew",
        "Item_Cook_K_09": "Milk",
        "Item_Material_05_00": "Milk x0",
        "Item_Cook_L_01": "Monster Stew",
        "Item_Cook_L_02": "Monster Soup",
        "Item_Cook_L_03": "Monster Cake",
        "Item_Cook_L_04": "Monster Rice Balls",
        "Item_Cook_L_05": "Monster Curry",
        "Item_Cook_M_01": "Wheat Bread",
        "Item_Cook_N_01": "Seafood Paella",
        "Item_Cook_N_02": "Fruitcake",
        "Item_Cook_N_03": "Vegetable Omelet",
        "Item_Cook_N_04": "Mushroom Omelet",
        "Item_Cook_O_01": "Dubious Food",
        "Item_Cook_O_02": "Rock-Hard Food",
        "Item_Cook_P_01": "Fragrant Mushroom Sauté",
        "Item_Cook_P_02": "Herb Sauté",
        "Item_Cook_P_03": "Spiced Meat Skewer",
        "Item_Cook_P_04": "Prime Spiced Meat Skewer",
        "Item_Cook_P_05": "Gourmet Spiced Meat Skewer"
      }
    },
    {
      "id": "other",
      "items": {
        "Obj_DungeonClearSeal": "Spirit Orb",
        "Obj_KorokNuts": "Korok Seed",
        "PlayerStole2": "Paraglider",
        "Obj_ProofBook": "Classified Envelope",
        "Obj_DRStone_Get": "Sheikah Slate",
        "Obj_HeroSoul_Zora": "Mipha's Grace",
        "Obj_HeroSoul_Gerudo": "Urbosa's Fury",
        "Obj_HeroSoul_Goron": "Daruk's Protection",
        "Obj_HeroSoul_Rito": "Revali's Gale",
        "Obj_DLC_HeroSoul_Zora": "Mipha's Grace + (v1.4)",
        "Obj_DLC_HeroSoul_Gerudo": "Urbosa's Fury + (v1.4)",
        "Obj_DLC_HeroSoul_Goron": "Daruk's Protection + (v1.4)",
        "Obj_DLC_HeroSoul_Rito": "Revali's Gale + (v1.4)",
        "Obj_Maracas": "Hestu's Maracas",
        "Obj_ProofKorok": "Hestu's Gift",
        "Obj_ProofSandwormKiller": "Medal of Honor: Molduga",
        "Obj_ProofGiantKiller": "Medal of Honor: Hinox",
        "Obj_ProofGolemKiller": "Medal of Honor: Talus",
        "KeySmall": "Small Key",
        "Obj_Armor_115_Head": "Thunder Helm",
        "GameRomHorseReins_00": "Stable Bridle",
        "GameRomHorseReins_01": "Traveler's Bridle",
        "GameRomHorseReins_02": "Royal Reins",
        "GameRomHorseReins_03": "Knight's Bridle",
        "GameRomHorseReins_04": "Monster Bridle",
        "GameRomHorseReins_05": "Extravagant Bridle",
        "GameRomHorseReins_10": "Ancient Bridle (v1.4)",
        "GameRomHorseSaddle_00": "Stable Saddle",
        "GameRomHorseSaddle_01": "Traveler's Saddle",
        "GameRomHorseSaddle_02": "Royal Saddle",
        "GameRomHorseSaddle_03": "Knight's Saddle",
        "GameRomHorseSaddle_04": "Monster Saddle",
        "GameRomHorseSaddle_05": "Extravagant Saddle",
        "GameRomHorseSaddle_10": "Ancient Saddle (v1.4)",
        "GameRomHorse00S": "Donkey",
        "GameRomHorseNushi": "Lord of the Mountain",
        "GameRomHorseEpona": "Epona",
        "GameRomHorseBone": "Stalhorse",
        "GameRomHorse00L": "Giant Horse",
        "GameRomHorseZelda": "Royal White Stallion",
        "Obj_WarpDLC": "Travel Medallion (v1.3)",
        "Get_TwnObj_DLC_MemorialPicture_A_01": "Picture of the Champions (v1.4)"
      }
    }
  ]
}
`

// InventoryItem represents an item in the BOTW inventory
type InventoryListItem struct {
	Name     string // Display name
	RawName  string // Original item ID
	Quantity uint32 // Item quantity or durability
	Index    int    // Position in the inventory array
	Category string // Item category
}

// FilterValue implements list.Item interface
func (i InventoryListItem) FilterValue() string { return i.Name }

// InventoryItemDelegate is a delegate for rendering inventory items
type InventoryItemDelegate struct{}

// Height implements list.ItemDelegate
func (d InventoryItemDelegate) Height() int { return 1 }

// Spacing implements list.ItemDelegate
func (d InventoryItemDelegate) Spacing() int { return 0 }

// Update implements list.ItemDelegate
func (d InventoryItemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

// Render implements list.ItemDelegate
func (d InventoryItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(InventoryListItem)
	if !ok {
		return
	}

	// Define styles
	selected := m.Index() == index
	titleStyle := ItemTitleStyle
	catStyle := CategoryStyle
	quantityStyle := QuantityStyle

	if selected {
		titleStyle = titleStyle.Foreground(lipgloss.Color("#ffffff")).Background(lipgloss.Color("#0a7bff"))
		catStyle = catStyle.Foreground(lipgloss.Color("#ffffff")).Background(lipgloss.Color("#0a7bff"))
		quantityStyle = quantityStyle.Foreground(lipgloss.Color("#ffffff")).Background(lipgloss.Color("#0a7bff"))
	}

	// Prepare value display based on category
	valueText := fmt.Sprintf("%d", i.Quantity)
	if i.Category == "Weapons" || i.Category == "Bows" || i.Category == "Shields" {
		valueText = fmt.Sprintf("Durability: %d", i.Quantity)
	} else if i.Quantity > 1 {
		valueText = fmt.Sprintf("x%d", i.Quantity)
	} else if i.Category == "Key Items" {
		valueText = "" // Don't show quantity for key items
	}

	// Format and write
	title := titleStyle.Render(i.Name)
	category := catStyle.Render(fmt.Sprintf("[%s]", i.Category))
	quantity := quantityStyle.Render(valueText)

	row := lipgloss.JoinHorizontal(
		lipgloss.Left,
		title,
		lipgloss.NewStyle().Width(3).Render(""), // Spacer
		category,
		lipgloss.NewStyle().Width(3).Render(""), // Spacer
		quantity,
	)

	fmt.Fprint(w, row)
}

// ItemTitleStyle is the style for item titles
var ItemTitleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#ffffff"))

// CategoryStyle is the style for item categories
var CategoryStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#8a8a8a"))

// QuantityStyle is the style for item quantities
var QuantityStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00cc00"))

// InventorySection contains information about where inventory data is stored
type InventorySection struct {
	ItemsHash         uint32              // Hash for ITEMS section (0x5f283289)
	ItemsQuantityHash uint32              // Hash for ITEMS_QUANTITY section (0x6a09fc59)
	ItemsOffset       int                 // Calculated offset for ITEMS section
	QuantitiesOffset  int                 // Calculated offset for ITEMS_QUANTITY section
	Items             []InventoryListItem // Parsed inventory items
	CurrentCategory   string              // Current filter category
	Categories        []string            // Available item categories
}

// TranslationCategory represents a category of item translations
type TranslationCategory struct {
	ID    string            `json:"id"`
	Items map[string]string `json:"items"`
}

// Translations contains all the translation categories
type Translations struct {
	Categories []TranslationCategory `json:"categories"`
}

// All translations loaded from JSON
var (
	botwTranslations   Translations
	translationsLoaded bool = false
)

// LoadTranslations loads the translations from the JSON file
func LoadTranslations() error {
	// Check if translations are already loaded
	if translationsLoaded {
		return nil
	}

	// unmarshall the json data from translationsRaw

	unmarshalledData := []byte(translationsRaw)
	if err := json.Unmarshal(unmarshalledData, &botwTranslations); err != nil {
		return fmt.Errorf("error parsing translations data: %v", err)
	}

	// Mark translations as Loaded
	translationsLoaded = true

	log.Printf("Loaded %d translation categories", len(botwTranslations.Categories))
	return nil

	/* // Read the translations file
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading translations file: %v", err)
	}

	// Parse the JSON
	err = json.Unmarshal(data, &botwTranslations)
	if err != nil {
		return fmt.Errorf("error parsing translations file: %v", err)
	}

	translationsLoaded = true
	log.Printf("Loaded %d translation categories", len(botwTranslations.Categories))
	return nil */
}

// GetItemDisplayName converts raw item names to human-readable names
func GetItemDisplayName(itemName string) string {
	// Try to load translations if not already loaded
	if !translationsLoaded {
		err := LoadTranslations()
		if err != nil {
			log.Printf("Warning: Could not load translations: %v", err)
		}
	}

	// Check if translations are loaded and search in all categories
	if translationsLoaded {
		for _, category := range botwTranslations.Categories {
			if name, ok := category.Items[itemName]; ok {
				return name
			}
		}
	}

	// If no mapping exists or translations aren't loaded, clean up the raw name
	name := strings.TrimPrefix(itemName, "Weapon_")
	name = strings.TrimPrefix(itemName, "Item_")
	name = strings.TrimPrefix(name, "Obj_")
	name = strings.Replace(name, "_", " ", -1)
	return name
}

// GetItemCategory determines the category of an item based on its name
func GetItemCategory(itemName string) string {
	// Try to load translations if not already loaded
	if !translationsLoaded {
		err := LoadTranslations()
		if err != nil {
			log.Printf("Warning: Could not load translations: %v", err)
		}
	}

	// If translations are loaded, check if item exists in any category
	if translationsLoaded {
		for _, category := range botwTranslations.Categories {
			for id := range category.Items {
				if id == itemName {
					return category.ID
				}
			}
		}
	}

	// Fallback to pattern matching if not found in translations
	if strings.HasPrefix(itemName, "Weapon_Sword") || strings.HasPrefix(itemName, "Weapon_Lsword") || strings.HasPrefix(itemName, "Weapon_Spear") {
		return "weapons"
	} else if strings.HasPrefix(itemName, "Weapon_Bow") || strings.Contains(itemName, "Arrow") {
		return "bows"
	} else if strings.HasPrefix(itemName, "Weapon_Shield") {
		return "shields"
	} else if strings.HasPrefix(itemName, "Armor_") {
		return "clothes"
	} else if strings.HasPrefix(itemName, "Item_Material") || strings.HasPrefix(itemName, "Item_Enemy_") ||
		strings.HasPrefix(itemName, "Item_Ore_") || strings.HasPrefix(itemName, "Item_Fruit_") ||
		strings.HasPrefix(itemName, "Item_Plant_") || strings.HasPrefix(itemName, "Animal_Insect_") {
		return "materials"
	} else if strings.HasPrefix(itemName, "Item_Roast_") || strings.HasPrefix(itemName, "Item_Cook_") ||
		strings.HasPrefix(itemName, "Item_Chilled") || strings.HasPrefix(itemName, "Item_RoastFish") ||
		strings.HasPrefix(itemName, "Item_Boiled") {
		return "food"
	} else if strings.HasPrefix(itemName, "Obj_") || strings.HasPrefix(itemName, "PlayerStole") ||
		strings.HasPrefix(itemName, "GameRom") {
		return "other"
	}
	return "other"
}

{-# LANGUAGE TemplateHaskell #-}

module Person
  ( Address (..)
  , Born    (..)
  , Name    (..)
  , Person  (..)
  , bornStreet
  , renameStreets
  , setBirthMonth
  , setCurrentStreet
  ) where

import Data.Time.Calendar (Day, toGregorian, fromGregorian)
import Control.Lens

data Name
  = Name { _foreNames :: String, _surName :: String }
makeLenses ''Name

data Address
  = Address { _street :: String, _houseNumber :: Int, _place :: String, _country :: String }
makeLenses ''Address

data Born
  = Born { _bornAt :: Address, _bornOn :: Day }
makeLenses ''Born

data Person
  = Person { _name :: Name, _born :: Born, _address :: Address }
makeLenses ''Person

bornStreet :: Born -> String
bornStreet = view (bornAt . street)

setCurrentStreet :: String -> Person -> Person
setCurrentStreet = set (address . street)

setBirthMonth :: Int -> Person -> Person
setBirthMonth month = over (born . bornOn) (setMonth month)
  where
    setMonth month day
      = let (year, _, date) = toGregorian day
        in fromGregorian year month date

renameStreets :: (String -> String) -> Person -> Person
renameStreets f = (over (address . street) f) . (over (born . bornAt . street)) f
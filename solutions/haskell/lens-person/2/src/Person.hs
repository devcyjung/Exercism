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
bornStreet = (^. bornAt . street)

setCurrentStreet :: String -> Person -> Person
setCurrentStreet s = address . street .~ s

setBirthMonth :: Int -> Person -> Person
setBirthMonth m = born . bornOn . getYMD . _2 .~ m
  where
    getYMD = iso toGregorian (\(y, m, d) -> fromGregorian y m d)

renameStreets :: (String -> String) -> Person -> Person
renameStreets f = (born . bornAt . street %~ f) . (address . street %~ f)
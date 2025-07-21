module Port
  IDENTIFIER = :PALE

  def self.get_identifier(city)
    city[...4].upcase.to_sym
  end
  
  def self.get_terminal(ship_identifier)
    case ship_identifier.to_s[...3]
      in "OIL" | "GAS"
        :A
      else
        :B
    end
  end
end

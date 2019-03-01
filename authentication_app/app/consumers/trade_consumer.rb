class TradeConsumer < Racecar::Consumer
  subscribes_to "rails_auth"

  def process(message)
    puts "Rails got Go's message from Kafka: { #{ message.key }: #{ message.value } }"
  end
end

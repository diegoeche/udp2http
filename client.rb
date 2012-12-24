require "socket"
class UDPRequest
  def udp_socket
    @udp_socket ||= UDPSocket.new
  end

  def request(method, path, body)
    proxy_host  = "localhost"
    proxy_port  = 2112
    method      = method_flag( method )
    path_length = path.length
    body_length = body == nil ? 0 : body.length

    # <<Method:16/big,PathLength:16/big,Path:PathLength/binary,BodyLength:16/big,Body:BodyLength/binary>>
    packet = [method,
              path_length,
              path, body_length,
              body].pack("nnA*#{path_length}nA*#{body_length}")
    p packet
    udp_socket.send( packet, 0, proxy_host, proxy_port )
    udp_socket.flush
  rescue Exception => e
    if udp_socket.closed?
      @udp_socket = UDPSocket.new
    end
    false
  end

  def post(path, req_body)
    self.request(:post, path, req_body)
  end

  def method_flag( method )
    case method
    when :get
      0
    when :post
      1
    when :put
      2
    when :delete
      3
    end
  end
end

# Test
puts "hello"
client = UDPRequest.new
client.post("http://graph.facebook.com", '"This is the Body"')

class UsersController < ApplicationController
  before_action :set_user, only: [ :edit, :update, :show, :destroy ]

  def create
    @user = User.new( user_params )

    if @user.save
      DeliveryBoy.deliver_async( @user.to_json, topic: "apex" )
      render json: @user
    else
      render json: @user.errors
    end
  end

  def show
  end

  private

  def user_params
    params.require( :user ).permit( :full_name, :email, :password, :password_confirmation )
  end

  def set_user
    @user = User.find( params[ :id ] )
  end
end
